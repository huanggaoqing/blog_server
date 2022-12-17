package controller

import (
	"blog_server/dto/request"
	"blog_server/dto/response"
	"blog_server/resp"
	"blog_server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleController struct {
	service *service.ArticleService
}

func (a *ArticleController) Save(ctx *gin.Context) {
	params := &request.SaveArticleRequest{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.Error(resp.PARAMS_ERROR)
		return
	}
	id, err := a.service.Save(params)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp.NewResponse[string](id))
}

func (a *ArticleController) GetList(ctx *gin.Context) {
	params := &request.GetArticleListRequest{}
	err := ctx.ShouldBind(params)
	if err != nil {
		ctx.Error(resp.PARAMS_ERROR)
		return
	}
	articleList, err := a.service.GetList(params)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp.NewResponse[response.GetArticleListResponseByPaging](articleList))
}

func (a *ArticleController) GetDetail(ctx *gin.Context) {
	params := &request.GetArticleDetailRequest{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.Error(resp.PARAMS_ERROR)
		return
	}
	articleDetail, err := a.service.GetDetailByArticleId(params.ArticleId)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp.NewResponse[*response.GetArticleDetailResponse](articleDetail))
}

func NewArticleController() *ArticleController {
	return &ArticleController{
		service: service.NewArticleService(),
	}
}
