package main

import (
	"blog_server/db"
	"blog_server/logger"
	"blog_server/router"
	"blog_server/tools"
	"log"
)

// 初始化一些所需资源
func init() {
	_, err := tools.InitConfig()
	if err != nil {
		log.Panicln("初始化程序配置失败%", err)
	}
	err = logger.InitLogger()
	if err != nil {
		log.Panicln("日志初始化失败", err)
	}
	_, err = db.InitDb()
	if err != nil {
		log.Panicln("初始化数据库失败%", err)
	}
}

// @Title blog_server
// @Version v1
// @Description 个人网站服务端
// @Contact.name huanggaoqing
// @Contact.url https://github.com/huanggaoqing
// @BasePath /
func main() {
	r := router.RegisterRouter()
	serverConfig := tools.GetSysConfig().Server
	err := r.Run(serverConfig.Host + serverConfig.Port)
	if err != nil {
		logger.Log.Errorf("开启服务失败%+v", err)
	}
}
