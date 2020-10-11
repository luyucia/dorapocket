package main

import (
	"fmt"
	"github.com/gohouse/converter"
	"github.com/labstack/echo/v4"
)

func main() {
	//实例化echo对象。
	e := echo.New()

	InitRouter(e)

	err := converter.NewTable2Struct().
		SavePath("h:/dorapocket/model/model.go").
		Dsn("root:MhxzKhl*@tcp(101.200.46.254:3306)/wenwen?charset=utf8").
		EnableJsonTag(true).
		TagKey("gorm").
		Run()
	fmt.Println(err)

	e.Start(":80")
}
