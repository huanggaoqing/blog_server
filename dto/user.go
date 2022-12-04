package dto

type UserRegisterRequest struct {
	UserName string `json:"userName" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
	Phone    string `json:"phone" binding:"required"`    // 手机号
}

type UserLoginRequest struct {
	Type     string `json:"type" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	UserInfo *ResponseUserInfo `json:"userInfo"`
	Token    string            `json:"token"`
}

type ResponseUserInfo struct {
	UserId      int    `json:"userId"`
	UserName    string `json:"userName"`
	PhoneNumber string `json:"phoneNumber"`
	Avatar      string `json:"avatar"`
}
