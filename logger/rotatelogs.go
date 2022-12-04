package logger

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"time"
)

type rotateLogsParams struct {
	filename     string
	maxDay       time.Duration
	rotationTime time.Duration
}

// createRotateLogsConfig 创建日志切割实例
func createRotateLogsConfig(params *rotateLogsParams) (*rotatelogs.RotateLogs, error) {
	l, err := rotatelogs.New(
		params.filename+".%Y%m%d%H%M.log",
		rotatelogs.WithMaxAge(params.maxDay),
		rotatelogs.WithRotationTime(params.rotationTime),
	)
	if err != nil {
		return nil, err
	}
	return l, nil
}
