//+build ignore

package main

import (
	"fmt"
	"github.com/gohouse/converter"
)
//go:generate go run generate_models.go
func main() {
	//实例化echo对象。

	err := converter.NewTable2Struct().
		SavePath("model/model.go").
		Dsn("root:MhxzKhl*@tcp(101.200.46.254:3306)/wenwen?charset=utf8").
		EnableJsonTag(true).
		TagKey("gorm").
		Run()
	fmt.Println(err)

}
