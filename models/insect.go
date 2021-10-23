package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//insect : ""
type Insect struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Activestatus bool               `json:"activestatus" form:"activestatus" bson:"activestatus,omitempty"`
	Name         string             `json:"name" form:"name" bson:"name,omitempty"`
	Status       string             `json:"status" form:"status" bson:"status,omitempty"`
	Created      *Created           `json:"created" form:"created" bson:"created,omitempty"`
	Version      int                `json:"version" form:"version" bson:"version,omitempty"`
}
type InsectFilter struct {
	OmitIDs      []primitive.ObjectID `json:"omitIds,omitempty"`
	ActiveStatus []bool               `json:"activestatus,omitempty"`
	Status       []string             `json:"status" form:"status" bson:"status,omitempty"`
	SortBy       string               `json:"sortBy"`
	SortOrder    int                  `json:"sortOrder"`
	Searchbox    struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}
type RefInsect struct {
	Insect `bson:",inline"`
	Ref    struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
