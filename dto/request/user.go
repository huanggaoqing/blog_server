package request

type UserRegisterRequest struct {
	UserName string `json:"userName" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
	Phone    string `json:"phone" binding:"required"`    // 手机号
}

type UserLoginRequest struct {
	Type     string `json:"type"  :"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}
