package dao

import (
	"context"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoBaseDao struct {
	Collection *mongo.Collection
}

func NewBaseDao(collection *mongo.Collection) MongoBaseDao {
	return MongoBaseDao{collection}
}

func (b *MongoBaseDao) Insert(ctx context.Context, data interface{}) (primitive.ObjectID,error) {

	insertResult, err := b.Collection.InsertOne(ctx, data)
	ObjectId := insertResult.InsertedID.(primitive.ObjectID)

	if err != nil {
		log.Error(err)
		return primitive.ObjectID{},err
	}

	return ObjectId,err
}


func (b *MongoBaseDao) Update(ctx context.Context,filter interface{},update interface{}) (int64,error) {

	result, err := b.Collection.UpdateOne(ctx,filter,update)

	if err != nil {
		log.Error(err)
		return  0,err
	}
	return result.ModifiedCount,err
}

func (b *MongoBaseDao) Delete(ctx context.Context,filter interface{}) (int64,error) {

	result, err := b.Collection.DeleteOne(ctx,filter)

	if err != nil {
		log.Error(err)
		return 0,err
	}
	return result.DeletedCount,err
}



func (b *MongoBaseDao) QueryOne(ctx context.Context,filter bson.M, s interface{}) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := b.Collection.FindOne(ctx, filter).Decode(s)

	if err != nil {
		log.Error(err)

	}
}

func (b *MongoBaseDao) QueryList(ctx context.Context,filter bson.M, s interface{}) []interface{}{

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	cur, err := b.Collection.Find(ctx, filter)

	if err != nil {
		log.Error(err)
	}

	defer cur.Close(ctx)
	var resultList []interface{}
	for cur.Next(ctx) {
		var t = s
		err := cur.Decode(&t)
		if err != nil {
			log.Error(err)
		}
		resultList = append(resultList,t)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return resultList

}
