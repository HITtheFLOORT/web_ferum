package Middleware

import "github.com/gin-gonic/gin"

func MiddlewareRegist()[]func(c *gin.Context){
	var a=[]func(c *gin.Context){}
	a=append(a, JWTAuthMiddleware())
	return a
}
