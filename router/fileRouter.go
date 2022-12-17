package router

import (
	"blog_server/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func fileRouter(r *gin.RouterGroup) error {
	file := r.Group("file")
	absPath, err := tools.GetAbsPath("image")
	if err != nil {
		return err
	}
	file.StaticFS("/image", http.Dir(absPath))
	return nil
}
