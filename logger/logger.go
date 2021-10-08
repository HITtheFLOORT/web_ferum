/*
zap
将日志写入文件而不是终端
修改时间编码器
在日志文件中使用大写字母记录日志级别
Lumberjack
为了添加日志切割归档功能，我们将使用第三方库Lumberjack来实现。
日志文件每1MB会切割并且在当前目录下最多保存5个备份
gin-logger/recovery
我们可以模仿Logger()和Recovery()的实现，使用我们的日志库来接收gin框架默认输出的日志。
 */
package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)
func Init()(err error){
	writer:=getLogWriter()
	encode:=getEncoder()
	var l=new(zapcore.Level)
	if err=l.UnmarshalText([]byte(viper.GetString("log.level")));err!=nil{
		fmt.Printf("get zap setting failed,err:%s\n",err.Error())
		return
	}
	if viper.GetString("app.mode")=="dev"{
		consoleencode:=zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core:=zapcore.NewTee(
			zapcore.NewCore(consoleencode,zapcore.Lock(os.Stdout),zapcore.DebugLevel),
			zapcore.NewCore(encode,writer,l),
			)
		lg:=zap.New(core,zap.AddCaller())
		zap.ReplaceGlobals(lg)
	}else{

		core:=zapcore.NewCore(encode,writer,l)
		lg:=zap.New(core,zap.AddCaller())
		zap.ReplaceGlobals(lg)
	}

	return
}
func getEncoder()zapcore.Encoder{
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   viper.GetString("log.filename"),
		MaxSize:    viper.GetInt("log.MaxSize"),
		MaxBackups: viper.GetInt("log.MaxBackups"),
		MaxAge:     viper.GetInt("log.MaxAge"),
		Compress:   viper.GetBool("log.Compress"),
	}
	return zapcore.AddSync(lumberJackLogger)
}


func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
