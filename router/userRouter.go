package router

import (
	"blog_server/controller"
	"github.com/gin-gonic/gin"
)

// registerUserRouter 注册用户模块路由
func registerUserRouter(r *gin.RouterGroup) {
	userController := controller.NewUserController()
	user := r.Group("/user")
	{
		// 用户注册
		user.POST("/register", userController.Register)
		// 用户登录
		user.POST("/login", userController.Login)
	}
}
