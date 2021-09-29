package Config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.New()
	viper.SetConfigName("app")
	viper.AddConfigPath("./")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("未找到该文件")
		} else {
			panic("读取失败")
		}
	}
	//自动监听配置修改
	viper.WatchConfig()
}