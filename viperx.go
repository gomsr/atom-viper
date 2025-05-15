package viperx

import (
	"github.com/gomsr/atom-viper/initialize"
	"github.com/spf13/viper"
)

func InitViper(conf any, path ...string) *viper.Viper {
	// init viper
	return initialize.Viper(conf, path...) // 初始化Viper
}

//func InitViperWithConf(conf any, path ...string) *viper.Viper {
//	// init viper
//	v := initialize.Viper(&kg.C, path...) // 初始化Viper
//	initialize.Viper(conf, path...)       // 初始化Viper
//
//	return v
//}
