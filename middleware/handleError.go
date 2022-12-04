package middleware

import (
	"blog_server/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleError 统一错误处理
func HandleError(ctx *gin.Context) {
	ctx.Next()
	for _, e := range ctx.Errors {
		err := e.Err
		if _, ok := err.(*resp.MyError); ok {
			ctx.JSON(http.StatusOK, err)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "服务内部错误",
			})
		}
	}
}
