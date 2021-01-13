package main

import (
	"dorapocket/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	_ "dorapocket/docs"
)

func InitRouter(r *echo.Echo) {


	r.Use(middleware.CORS())

	r.GET("/ping", controllers.Ping)

	// 指定swagger自动文档路由
	r.GET("/docs/*swagger", echoSwagger.WrapHandler)




	// -------------项目管理----------
	// 创建项目
	r.POST("/project/create", controllers.CreateProject)
	// 项目列表
	r.GET("/project/list", controllers.ListProject)
	// 修改项目
	r.POST("/project/edit", controllers.UpdateProject)
	// 删除项目
	r.POST("/project/delete", controllers.DeleteProject)


}
