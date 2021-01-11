package service

import (
	"context"
	"dorapocket/dao"
	"dorapocket/model"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectService struct {
	dao dao.MongoBaseDao
}

func NewProjectService(dao dao.MongoBaseDao)  ProjectService{
	return ProjectService{dao:dao}
}

func (service ProjectService) Create(ctx context.Context,project *model.Project) error {
	objectId,err := service.dao.Insert(ctx,project)

	log.Info(objectId)

	return err

}

func (service ProjectService) Update(ctx context.Context,id string,project *model.Project) (int64,error) {

	ObjectId,err :=  primitive.ObjectIDFromHex(id)

	if err!=nil {
		log.Error(err)
	}

	filter := bson.D{
		{"_id",ObjectId},
	}

	update:= bson.M{
		"$set":project,
	}

	return service.dao.Update(ctx,filter,update)
}

// 物理删除
func (service ProjectService) CleanOne(ctx context.Context,id string) (int64,error) {

	ObjectId,err :=  primitive.ObjectIDFromHex(id)

	if err!=nil {
		log.Error(err)
		return 0,err
	}

	filter := bson.D{
		{"_id",ObjectId},
	}

	return service.dao.Delete(ctx,filter)
}

func (service ProjectService) List(ctx context.Context) []interface{} {
	//var projects []model.Project
	//projects = service.dao.QueryList(nil,model.Project{})

	return service.dao.QueryList(ctx,nil,model.Project{})
}