package global

import (
	"ShareDex-BE/src/constant"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// InitLogger 初始化日志器
func InitLogger() {
	var err error
	setZapEncoderConfig(&ZapLoggerConf.EncoderConfig)
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

func setZapEncoderConfig(encoderConfig *zapcore.EncoderConfig) {

	encoderConfig.EncodeTime = customTimeEncoder // 设置自定义的时间编码器
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.EncodeName = zapcore.FullNameEncoder
}

// 自定义时间编码器
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	// 设置时间的格式
	enc.AppendString(t.Format(constant.DateTimeMilli))
}
