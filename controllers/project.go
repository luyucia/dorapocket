package controllers

import (
	"dorapocket/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

// 创建项目
// @Summary 创建项目
// @Description 创建项目
// @Accept  json
// @Produce  json
// @Param name formData string false "项目名称"
// @Param description formData string false "项目描述"
// @Header 200 {string} Token " "
// @Router /project/create [post]
func CreateProject(c echo.Context) error {

	p := new(model.Project)

	if err := c.Bind(p); err != nil {
		return err
	}

	service := initializePorjectService()
	id, err := service.Create(c.Request().Context(), p)
	if err != nil {
		return c.JSON(http.StatusOK, NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, NewDataResponse(id))
}

// @Summary 项目列表
// @Description 项目列表
// @Accept  json
// @Produce  json
// @Header 200 {string} Token " "
// @Router /project/list [get]
func ListProject(c echo.Context) error {

	p := new(model.Project)

	if err := c.Bind(p); err != nil {
		return err
	}

	service := initializePorjectService()

	rs := service.List(c.Request().Context())
	return c.JSON(http.StatusOK, NewDataResponse(rs))
}

// @Summary 项目编辑
// @Description 项目编辑
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param id formData string false "ID"
// @Param name formData string false "项目名称"
// @Param description formData string false "项目描述"
// @Header 200 {string} Token " "
// @Router /project/edit [post]
func UpdateProject(c echo.Context) error {

	p := new(model.Project)

	if err := c.Bind(p); err != nil {
		return err
	}

	ObjectId := p.Id
	//p.Name = c.FormValue("name")
	//p.Description = c.FormValue("description")

	service := initializePorjectService()
	cs, err := service.Update(c.Request().Context(), ObjectId, p)
	if err != nil {
		return c.JSON(http.StatusOK, NewErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, NewDataResponse(cs))
}

// @Summary 项目删除
// @Description 项目删除
// @Accept  x-www-form-urlencoded
// @Produce  json
// @Param id formData string false "ID"
// @Header 200 {string} Token " "
// @Router /project/delete [post]
func DeleteProject(c echo.Context) error {

	p := new(model.Project)

	if err := c.Bind(p); err != nil {
		return err
	}

	ObjectId := p.Id

	service := initializePorjectService()

	rs, err := service.CleanOne(c.Request().Context(), ObjectId)

	if err != nil {
		return c.JSON(http.StatusOK, NewErrorResponse(err.Error()))
	}
	return c.JSON(http.StatusOK, NewDataResponse(rs))
}
