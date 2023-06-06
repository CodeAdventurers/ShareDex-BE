package global

import "go.uber.org/zap/zapcore"

var (
	ZapLoggerConf *ZapLoggerConfig = &ZapLoggerConfig{}
	MySQLConf     *MySQLConfig     = &MySQLConfig{}
	RedisConf     *RedisConfig     = &RedisConfig{}
)

type ZapLoggerConfig struct {
	Level            string                 `mapstructure:"level"`
	Encoding         string                 `mapstructure:"encoding"`
	OutputPaths      []string               `mapstructure:"outputPaths"`
	ErrorOutputPaths []string               `mapstructure:"errorOutputPaths"`
	InitialFields    map[string]interface{} `mapstructure:"initialFields"`
	EncoderConfig    zapcore.EncoderConfig  `mapstructure:"encoderConfig"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}
