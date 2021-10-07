package routers

import (
	"buble/dao/controller"
	"buble/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup()(r *gin.Engine){
	r=gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	fisrpage_router(r)
	signup_router(r)
	return r
}
func fisrpage_router(r *gin.Engine){
	r.GET("/index", func(c *gin.Context) {
		c.String(http.StatusOK,"first page")
	})
}
func signup_router(r *gin.Engine){
	r.POST("/signup", func(c *gin.Context) {
		controller.SignHandler(c)
	})
}