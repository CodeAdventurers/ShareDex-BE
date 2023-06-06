package global

import (
	"github.com/spf13/viper"
)

type viperManager struct {
}

// NewViperManger 构造
func NewViperManger() *viperManager {
	vm := &viperManager{}
	return vm
}

func (vm *viperManager) InitConfig() {
	// 加载日志配置
	vm.loadLoggerConfig()

	// mysql、redis、minio
}

func (vm *viperManager) loadLoggerConfig() {
	loggerViper := viper.New()
	loggerViper.SetConfigName("logger")       // 配置文件名称（无需后缀）
	loggerViper.SetConfigType("yaml")         // 配置文件类型
	loggerViper.AddConfigPath("./src/config") // 配置文件路径

	err := loggerViper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = loggerViper.Unmarshal(ZapLoggerConf)
	if err != nil {
		panic(err)
	}

}
