package global

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitLogger 初始化日志器
func InitLogger() {
	var err error
	cfg := zap.Config{
		Encoding:         ZapLoggerConf.Encoding,
		Level:            zap.NewAtomicLevelAt(getZapLevel(ZapLoggerConf.Level)),
		OutputPaths:      ZapLoggerConf.OutputPaths,
		ErrorOutputPaths: ZapLoggerConf.ErrorOutputPaths,
		EncoderConfig:    ZapLoggerConf.EncoderConfig,
	}

	Logger, err = cfg.Build()
	if err != nil {
		panic(err)
	}
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}
