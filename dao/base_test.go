
package dao

import (
	"dorapocket/core"
	"dorapocket/model"
	"go.mongodb.org/mongo-driver/bson"

	"testing"
)



func TestBaseDao_Insert(t *testing.T) {

	t.Log("插入测试")

	core.Init("../config/config.toml")

	project := model.Project{}
	project.Name = "测试项目1"

	dao,_ := initializeDao()

	//dao := MongoBaseDao{database:"dorabox"}

	dao.Insert(project)

	dao.QueryOne(bson.M{"name":"test"},&project)

	list := dao.QueryList(bson.M{"name":"测试项目9"},&project)

	t.Log(list)


	t.Log(project)



}
