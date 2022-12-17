package controller

import (
	"blog_server/dto/request"
	"blog_server/module/dbModule"
	"blog_server/resp"
	"blog_server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProfileController struct {
	service *service.ProfileService
}

func (p *ProfileController) Save(ctx *gin.Context) {
	params := &request.SaveProfileRequest{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.Error(resp.PARAMS_ERROR)
		return
	}
	profileId, err := p.service.Save(params)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp.NewResponse[int](profileId))
}

func (p *ProfileController) Get(ctx *gin.Context) {
	params := &request.GetProfileRequest{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.Error(resp.PARAMS_ERROR)
		return
	}
	profile, err := p.service.Get(params)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp.NewResponse[*dbModule.Profile](profile))
}

func (p *ProfileController) Set(ctx *gin.Context) {
	params := &request.SetProfileRequest{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.Error(resp.PARAMS_ERROR)
		return
	}
	profile, err := p.service.Set(params)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp.NewResponse[int](profile))
}

func NewProfileController() *ProfileController {
	return &ProfileController{
		service: service.NewProfileService(),
	}
}
