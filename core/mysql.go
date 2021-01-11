package core

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/gommon/log"
	"time"
)

var Db *gorm.DB

func InitMysql() {

	log.Info("init mysql ...")
	//配置读取
	username := Config.GetString("database.username")
	host := Config.GetString("database.host")
	password := Config.GetString("database.password")
	port := Config.GetString("database.port")
	database := Config.GetString("database.database")
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, host, port, database)
	//连接
	log.Debug(dsn)
	Db, err = gorm.Open("mysql", dsn)

	if err != nil {
		log.Error(err)
		panic("database connection failed")
	}

	//测试数据库连接
	err = Db.DB().Ping()
	if err != nil {
		log.Error(err)
		panic("database connection failed")
	}

	// 连接池设置
	Db.DB().SetMaxIdleConns(Config.GetInt("dbpool.max_idle_conn"))
	Db.DB().SetMaxOpenConns(Config.GetInt("dbpool.max_open_conn"))
	Db.DB().SetConnMaxLifetime(time.Minute * 5 )

	log.Info("init mysql done")

}
