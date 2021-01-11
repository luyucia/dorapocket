package core

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

var Config *viper.Viper

func InitConfig(confPath string)  {

	log.Info("init config ...")
	Config = viper.New()

	Config.SetConfigFile(confPath)
	Config.SetConfigType("toml")

	err  := Config.ReadInConfig()
	if err!=nil {
		log.Error(err)
	}

	log.Info("init config done")

}


