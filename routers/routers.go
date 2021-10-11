package routers

import (
	Middleware "buble/Middleware"
	"buble/controller"
	"buble/logger"
	"buble/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup()(r *gin.Engine){
	r=gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	utils.InitTrans("zh")
	fisrpage_router(r)
	signup_router(r)
	login_router(r)
	ping_router(r)
	_404_router(r)
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
func login_router(r *gin.Engine){
	r.POST("/login",func(c *gin.Context) {
		controller.LoginHandler(c)
	})
}
func ping_router(r *gin.Engine){
	r.POST("/ping", Middleware.JWTAuthMiddleware(),func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"success",
		})
	})
}
func _404_router(r *gin.Engine){
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound,gin.H{
			"msg":"pages not found",
		})
	})
}
