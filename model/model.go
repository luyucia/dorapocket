package model


type Project struct {

	Id string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" form:"name" query:"name"`
	Description string `json:"description" form:"description" query:"description"`

}