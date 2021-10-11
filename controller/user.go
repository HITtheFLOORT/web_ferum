package controller

import (
	"buble/logics"
	"buble/models"
	"buble/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
)

func SignHandler(c *gin.Context){
	//1.获取参数，参数检验
	var p =models.ParamSignUp{}
	if err:=c.ShouldBindJSON(&p);err!=nil{//json格式，字段类型
		zap.L().Error("Signup with invalid param",zap.Error(err))
		if errs,ok:=err.(validator.ValidationErrors);ok{
			c.JSON(http.StatusOK,gin.H{
				"msg":utils.RemoveTopStruct(errs.Translate(utils.Trans)),
			})
		}else{
			c.JSON(http.StatusOK,gin.H{
				"msg":err.Error(),
			})
		}
		return
	}
	//2.业务处理
	if err:=logics.Signup(&p);err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"msg":err.Error(),
		})
		return
	}
	//3.返回响应
	c.JSON(http.StatusOK,gin.H{
		"msg":"注册成功",
	})
	return
}
func LoginHandler(c *gin.Context){
	//1.获取参数
	var p=models.ParamLogin{}
	if err:=c.ShouldBindJSON(&p);err!=nil{
		zap.L().Error("Login with invalid param",zap.Error(err))
		if errs,ok:=err.(validator.ValidationErrors);ok{
			c.JSON(http.StatusOK,gin.H{
				"msg":utils.RemoveTopStruct(errs.Translate(utils.Trans)),
			})
		}else{
			c.JSON(http.StatusOK,gin.H{
				"msg":err.Error(),
			})
		}
		return
	}
	//2.业务处理
	token,err:=logics.Login(&p)
	if err!=nil{
		c.JSON(http.StatusOK,gin.H{
			"msg":err.Error(),
		})
		return
	}
	//3.返回响应
	c.JSON(http.StatusOK,gin.H{
		"msg":"登陆成功",
		"token":token,
	})
	return
}