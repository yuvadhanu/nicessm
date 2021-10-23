package models

//Language : ""
type Language struct {
	Name    string  `json:"name,omitempty" form:"name" bson:"name,omitempty"`
	Code    string  `json:"code,omitempty" form:"code" bson:"code,omitempty"`
	Status  string  `json:"status,omitempty" form:"status" bson:"status,omitempty"`
	Created Created `json:"created,omitempty" form:"created" bson:"created,omitempty"`
}
