package core

import "github.com/labstack/gommon/log"

func Init(confPath string) {

	log.Info("system init...")

	InitConfig(confPath)

	//InitMysql()

	InitMongo()

	log.Info("init success")
}
