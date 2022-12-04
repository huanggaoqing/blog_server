package router

import (
	"blog_server/controller"
	"github.com/gin-gonic/gin"
)

func profileRouter(r *gin.RouterGroup) {
	profileController := controller.NewProfileController()
	profile := r.Group("/profile")
	{
		profile.POST("/save", profileController.Save)
		profile.GET("/get", profileController.Get)
		profile.PUT("/set", profileController.Set)
	}
}
