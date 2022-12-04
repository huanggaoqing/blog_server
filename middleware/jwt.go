package middleware

import (
	"blog_server/constant"
	"blog_server/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

func JWT(ctx *gin.Context) {
	path := ctx.FullPath()
	clientType := ctx.GetHeader("clientType")
	if clientType == constant.CLIENT {
		constant.BaseWhiteList = append(constant.BaseWhiteList, constant.ClientWhiteLis...)
	}
	sort.Strings(constant.BaseWhiteList)
	idx := sort.SearchStrings(constant.BaseWhiteList, path)
	if idx < len(constant.BaseWhiteList) && constant.BaseWhiteList[idx] == path {
		ctx.Next()
		return
	}
	token := ctx.GetHeader("Authorization")
	_, err := tools.VerificationToken(token)
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		ctx.Abort()
	} else {
		ctx.Next()
	}
}
