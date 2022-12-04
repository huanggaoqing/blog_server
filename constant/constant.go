package constant

const (
	ADMIN  = "admin"
	CLIENT = "client"
)

var BaseWhiteList = []string{
	"/v1/user/login",
	"/v1/user/register",
}

var ClientWhiteLis = []string{
	"/v1/profile/get",
}
