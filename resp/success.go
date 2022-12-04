package resp

type response[T interface{}] struct {
	Code int    `json:"code"`
	Data T      `json:"data"`
	Msg  string `json:"msg"`
}

func NewResponse[T interface{}](data T) *response[T] {
	return &response[T]{
		Code: 200,
		Data: data,
		Msg:  "成功",
	}
}
