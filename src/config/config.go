package config

import "github.com/spf13/viper"

var Config = viper.New()

func InitConfig() {
	InitConfigByPath("./")
}

func InitConfigByPath(path string) {
	Config.AddConfigPath(path)          //设置配置文件路径
	Config.SetConfigName("application") //设置配置文件名
	Config.SetConfigType("yaml")        //设置配置文件类型
	//读取配置
	if err := Config.ReadInConfig(); err != nil {
		panic(err)
	}
}
