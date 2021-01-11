package core

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var Mongo *mongo.Client

func InitMongo()  {

	mongoUri := Config.GetString("mongo.uri")
	max_pool_size := Config.GetUint64("mongo.max_pool_size")
	min_pool_size := Config.GetUint64("mongo.min_pool_size")
	max_idle_time := Config.GetInt("mongo.max_idle_time")

	log.Info(mongoUri)
	log.Info("mongo.max_pool_size: ",max_pool_size)
	log.Info("mongo.min_pool_size: ",min_pool_size)
	log.Info("mongo.max_idle_time: ",max_idle_time)

	//mongo连接设置
	mongoConnectionOptions := options.Client().ApplyURI(mongoUri)
	mongoConnectionOptions.SetMaxPoolSize(max_pool_size)
	mongoConnectionOptions.SetMinPoolSize(min_pool_size)
	mongoConnectionOptions.SetMaxConnIdleTime(time.Duration(max_idle_time) * time.Second)

	client, err := mongo.NewClient(mongoConnectionOptions)

	if err!=nil{
		log.Error(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err!= nil {
		log.Error(err)
	}

	err = client.Ping(ctx, readpref.Primary())

	if err!= nil {
		log.Error(err)
	}

	Mongo = client
}


