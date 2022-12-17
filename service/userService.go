package service

import (
	"blog_server/dao"
	"blog_server/dto/request"
	"blog_server/dto/response"
	"blog_server/resp"
	"blog_server/tools"
)

type UserService struct {
	tokenInstance *tools.TokenCore
}

func (u *UserService) Register(params *request.UserRegisterRequest) (int, error) {
	userDao := dao.NewUserDao()
	if user, err := userDao.FindUserByPhone(params.Phone); err != nil {
		return -1, err
	} else if user != nil {
		return -1, resp.USER_EXISTS
	}
	params.Password = tools.Md5(params.Password)
	userId, err := userDao.Create(params)
	if err != nil {
		return -1, err
	}
	return userId, nil
}

func (u *UserService) Login(params *request.UserLoginRequest) (*response.UserLoginResponse, error) {
	userDao := dao.NewUserDao()
	params.Password = tools.Md5(params.Password)
	user, err := userDao.FindUserByPassword(params)
	if err != nil {
		return nil, err
	}
	// TODO 将token存入redis
	token, err := u.tokenInstance.GenerateToken(user.UserId, user.UserName)
	userData := &response.UserLoginResponse{
		UserInfo: &response.ResponseUserInfo{
			UserId:      user.UserId,
			UserName:    user.UserName,
			PhoneNumber: user.PhoneNumber,
			Avatar:      user.Avatar,
		},
		Token: token,
	}
	return userData, err
}

func NewUserService() *UserService {
	return &UserService{
		tokenInstance: tools.NewToken(),
	}
}
