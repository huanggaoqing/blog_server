package router

import (
	"blog_server/middleware"
	"blog_server/swagger"
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
func RegisterRouter() *gin.Engine {
	r := gin.Default()
	swagger.InitSwagger(r)
	r.Use(middleware.HandleError)
	r.Use(middleware.JWT())
	v1 := r.Group("v1")
	{
		registerUserRouter(v1)
		profileRouter(v1)
	}
	return r
}
