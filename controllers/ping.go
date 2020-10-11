package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// 创建问题
// @Summary 创建问题
// @Description 创建问题
// @Accept  json
// @Produce  json
// @Param title query string false "标题"
// @Param content query string false "内容"
// @Header 200 {string} Token " "
// @Router /question/create [post]
func Ping(c echo.Context) error {

	return c.String(http.StatusOK, "pong")
}
