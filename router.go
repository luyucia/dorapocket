package main

import (
	"dorapocket/controllers"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

func InitRouter(r *echo.Echo) {

	r.GET("/ping", controllers.Ping)

	// 指定swagger自动文档路由
	r.GET("/docs/*swagger", echoSwagger.WrapHandler)

	r.POST("/question/create", controllers.Ping)
	// 问题列表
	// 我的提问
	// 我的回答
	// 待回答问题
	// 回答问题
	// 回答邀请
	// 消息列表
	// 点赞
	// 收藏
	// 登录

}
