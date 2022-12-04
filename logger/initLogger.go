package logger

import (
	"blog_server/tools"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

var Log *zap.SugaredLogger

// InitLogger 初始化日志配置
func InitLogger() error {
	encoder := getEncoder()
	debuggerLogWrite, err := getLogWriter("info")
	errorLogWrite, err := getLogWriter("error")
	if err != nil {
		return err
	}
	var core zapcore.Core
	mode := tools.GetSysConfig().Mode
	if mode == "dev" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel)
	} else {
		debuggerCore := zapcore.NewCore(encoder, debuggerLogWrite, zapcore.DebugLevel)
		errorCore := zapcore.NewCore(encoder, errorLogWrite, zapcore.ErrorLevel)
		core = zapcore.NewTee(debuggerCore, errorCore)
	}
	Log = zap.New(core, zap.AddCaller()).Sugar()
	return nil
}

// getEncoder 获取日志输出编码格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// getLogWriter 获取日志输出文件位置
func getLogWriter(t string) (zapcore.WriteSyncer, error) {
	// 获取绝对路径
	dir := ""
	if t == "info" {
		dir = "log"
	} else {
		dir = "errorLog"
	}
	absPath, err := tools.GetAbsPath(dir)
	if err != nil {
		return nil, err
	}
	// 生成日志切割配置
	params := &rotateLogsParams{
		filename:     filepath.Join(absPath, "blog"),
		maxDay:       30 * 24 * time.Hour,
		rotationTime: time.Hour * 24,
	}
	l, err := createRotateLogsConfig(params)
	if err != nil {
		return nil, err
	}
	return zapcore.AddSync(l), nil
}
