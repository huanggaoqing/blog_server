package tools

import (
	"blog_server/module/configModule"
	"gopkg.in/yaml.v2"
	"os"
)

var sysConfig = &configModule.Module{}

// InitConfig 初始化程序配置
func InitConfig() (*configModule.Module, error) {
	// 获取配置文件路径
	p, err := GetAbsPath("config/config.yaml")
	if err != nil {
		return nil, err
	}
	// 打开配置文件
	file, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// 读取配置文件
	result, err := ReadAllByLine(file)
	// 将读取到的内容映射到module上
	if err := yaml.Unmarshal([]byte(result), sysConfig); err != nil {
		if err != nil {
			return nil, err
		}
	}
	return sysConfig, nil
}

func GetSysConfig() *configModule.Module {
	return sysConfig
}
