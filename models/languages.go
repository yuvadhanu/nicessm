package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Language : ""
type Languages struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name         string             `json:"name" form:"name" bson:"name,omitempty"`
	ActiveStatus bool               `json:"activestatus" form:"activestatus" bson:"activestatus,omitempty"`
	Status       string             `json:"status" form:"status" bson:"status,omitempty"`
	Created      *Created           `json:"created" form:"created" bson:"created,omitempty"`
	Version      string             `json:"version" form:"version" bson:"version,omitempty"`
}
type LanguageFilter struct {
	ActiveStatus []bool   `json:"activestatus,omitempty" form:"activestatus" bson:"activestatus,omitempty"`
	Status       []string `json:"status" form:"status" bson:"status,omitempty"`
	Version      []string `json:"version,omitempty" form:"version" bson:"version,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	Regex        struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
}
type RefLanguage struct {
	Languages `bson:",inline"`
	Ref       struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
