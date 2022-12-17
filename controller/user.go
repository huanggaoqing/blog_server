package controller

import (
	"blog_server/dto/request"
	"blog_server/dto/response"
	"blog_server/resp"
	"blog_server/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

// Register 用户注册
// @Summary 用户注册接口
// @Tags 用户接口
// @Params user body dto.UserRegisterRequest "参数"
// @Router /v1/user/register [post]
func (u *UserController) Register(ctx *gin.Context) {
	params := &request.UserRegisterRequest{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.Error(resp.PARAMS_ERROR)
		return
	}
	userId, err := u.service.Register(params)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp.NewResponse[int](userId))
}

func (u *UserController) Login(ctx *gin.Context) {
	params := &request.UserLoginRequest{}
	if err := ctx.ShouldBind(params); err != nil {
		ctx.Error(resp.PARAMS_ERROR)
		return
	}
	userInfo, err := u.service.Login(params)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, resp.NewResponse[*response.UserLoginResponse](userInfo))
}

func NewUserController() *UserController {
	return &UserController{
		service: service.NewUserService(),
	}
}
