package controller

import (
	"blog_server/dto"
	"blog_server/resp"
	"blog_server/service"
	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	service *service.ArticleService
}

func (a *ArticleController) Save(ctx *gin.Context) {
	params := &dto.SaveArticleRequest{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.Error(resp.PARAMS_ERROR)
		return
	}
	_, err := a.service.Save(params)
	if err != nil {
		ctx.Error(err)
		return
	}
}

func NewArticleController() *ArticleController {
	return &ArticleController{
		service: service.NewArticleService(),
	}
}
