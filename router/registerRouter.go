package router

import (
	"blog_server/logger"
	"blog_server/middleware"
	"blog_server/swagger"
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册路由
func RegisterRouter() *gin.Engine {
	r := gin.Default()
	swagger.InitSwagger(r)
	v1 := r.Group("v1")
	err := fileRouter(v1)
	if err != nil {
		logger.Log.Errorf("静态文件路由注册失败%s", err.Error())
	}
	v1.Use(middleware.HandleError)
	v1.Use(middleware.JWT())
	{
		registerUserRouter(v1)
		profileRouter(v1)
		articleRouter(v1)
	}
	return r
}
