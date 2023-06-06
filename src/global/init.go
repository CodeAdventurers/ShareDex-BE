package global

import (
	"go.uber.org/zap"
)

var (
	Logger *zap.Logger // 日志
)

func InitViper() {
	viperManager := NewViperManger()
	viperManager.InitConfig()
}
