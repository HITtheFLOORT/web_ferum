package controller

import (
	"buble/logics"
	"buble/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func SignHandler(c *gin.Context){
	//1.获取参数，参数检验
	 p :=models.ParamSignUp{}
	if err:=c.ShouldBindJSON(p);err!=nil{//json格式，字段类型
		zap.L().Error("Signup with invalid param",zap.Error(err))
		c.JSON(http.StatusOK,gin.H{
			"msg":"请求参数有误",
		})
		return
	}
	if len(p.Username)==0||len(p.Password)==0||p.RePassword!=p.Password{
		zap.L().Error("Signup with invalid param")
		c.JSON(http.StatusOK,gin.H{
			"msg":"请求参数有误",
		})
	}
	//2.业务处理
	logics.Signup()
	//3.返回响应
	c.JSON(http.StatusOK,gin.H{
		"msg":"success",
	})
}