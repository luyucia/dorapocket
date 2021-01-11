package dao

import (
	"dorapocket/core"
	"go.mongodb.org/mongo-driver/mongo"
)


func NewProjectCollection() *mongo.Collection {
	return core.Mongo.Database(core.Config.GetString("mongo.database")).Collection("projects")
}
