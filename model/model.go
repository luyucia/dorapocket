package model


type Project struct {

	Name string `json:"name" form:"name" query:"name"`
	Description string `json:"description" form:"description" query:"description"`

}