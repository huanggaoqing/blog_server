package router

import (
	"blog_server/controller"
	"github.com/gin-gonic/gin"
)

func articleRouter(r *gin.RouterGroup) {
	articleController := controller.NewArticleController()
	article := r.Group("/article")
	{
		article.POST("/save", articleController.Save)
		article.GET("/get", articleController.GetList)
		article.GET("/detail", articleController.GetDetail)
		article.DELETE("/delete")
		article.PUT("/update")
		article.PUT("/stick")
	}
}
