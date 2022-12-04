package resp

var (
	USER_EXISTS      = NewMyError(-1, "用户已存在")
	PARAMS_ERROR     = NewMyError(400, "参数错误")
	USER_NOT_EXISTS  = NewMyError(-1, "手机号或密码错误")
	NOT_PROFILE      = NewMyError(-1, "暂无简介")
	TOKEN_INVAILD    = NewMyError(401, "token无效")
	TOKEN_EXPIRED    = NewMyError(401, "token已失效")
	TOKEN_NOT_ACTIVE = NewMyError(401, "token未激活")
	NOT_TOKEN        = NewMyError(401, "token格式有误")
)

type MyError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *MyError) Error() string {
	return e.Msg
}

func NewMyError(code int, msg string) *MyError {
	return &MyError{
		Code: code,
		Msg:  msg,
	}
}
