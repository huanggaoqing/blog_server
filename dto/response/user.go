package response

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
