package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary ping
// @Description ping
// @Accept  json
// @Produce  json
// @Header 200 {string} Token " "
// @Router /question/create [post]
func Ping(c echo.Context) error {
	return c.JSON(http.StatusOK,NewSuccessResponse())
}

