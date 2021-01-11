package main

import (
	"dorapocket/core"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	//实例化echo对象。
	e := echo.New()

	InitRouter(e)

	core.Init("config/config.toml")

	port := core.Config.GetString("server.port")
	log.Info("server started on port ",port)

	e.Start(":"+port)
}
