// +build wireinject


package controllers

import (
	"dorapocket/dao"
	"dorapocket/service"
	"github.com/google/wire"
)
// 此处为依赖管理集合，负责注入依赖，使用google wire框架，执行wire命令 会生成依赖代码
func initializeDao() (dao.MongoBaseDao, error) {
	wire.Build(dao.NewBaseDao,dao.NewProjectCollection)
	return dao.MongoBaseDao{}, nil
}

func initializePorjectService() (service.ProjectService) {
	wire.Build(service.NewProjectService,dao.NewBaseDao ,dao.NewProjectCollection)
	return service.ProjectService{}
}
