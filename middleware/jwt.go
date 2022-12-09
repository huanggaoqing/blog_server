package middleware

import (
	"blog_server/constant"
	"blog_server/logger"
	"blog_server/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

func JWT() gin.HandlerFunc {
	tokenInstance := tools.NewToken()
	return func(ctx *gin.Context) {
		path := ctx.FullPath()
		clientType := ctx.GetHeader("clientType")
		if clientType == constant.CLIENT {
			constant.BaseWhiteList = append(constant.BaseWhiteList, constant.ClientWhiteList...)
		}
		sort.Strings(constant.BaseWhiteList)
		idx := sort.SearchStrings(constant.BaseWhiteList, path)
		if idx < len(constant.BaseWhiteList) && constant.BaseWhiteList[idx] == path {
			ctx.Next()
			return
		}
		token := ctx.GetHeader("Authorization")
		// TODO 在redis中查找用户对应的token是否存在
		_, err := tokenInstance.VerificationToken(token)
		if err != nil {
			logger.Log.Errorf("token校验失败：+%v", err)
			ctx.JSON(http.StatusOK, err)
			ctx.Abort()
		} else {
			ctx.Next()
		}
	}
}
