package constant

const (
	ADMIN  = "admin"
	CLIENT = "client"
)

const JwtSecret string = "myBlog"

var BaseWhiteList = []string{
	"/v1/user/login",
	"/v1/user/register",
}

var ClientWhiteList = []string{
	"/v1/profile/get",
}
