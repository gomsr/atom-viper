package viperx

import (
	"github.com/gomsr/atom-viper/initialize"
	"github.com/spf13/viper"
)

func InitViper(conf any, path ...string) *viper.Viper {
	// init viper
	return initialize.Viper(conf, path...) // 初始化Viper
}

func InitViperWithConf(conf1 any, conf2 any, path ...string) *viper.Viper {
	v := initialize.Viper(conf1, path...) // 初始化Viper
	initialize.Viper(conf2, path...)      // 初始化Viper

	return v
}
