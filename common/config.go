package common

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"youyin/model"
)

type Config struct {
	MYSQL model.MysqlConfig
	APP   model.APPConfig
	LOG   model.LogConfig
}

var Conf = &Config{}

func InitConfig(configPath string) *Config {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		err = viper.Unmarshal(&Conf)
		if err != nil {
			panic(err)
		}
	})
	return Conf

}
