package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//agroEcologicalZone : ""
type AgroEcologicalZone struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Activestatus bool               `json:"activestatus" form:"activestatus" bson:"activestatus,omitempty"`
	Name         string             `json:"name" form:"name" bson:"name,omitempty"`
	Status       string             `json:"status" form:"status" bson:"status,omitempty"`
	Created      *Created           `json:"created" form:"created" bson:"created,omitempty"`
	Version      int                `json:"version" form:"version" bson:"version,omitempty"`
	Zone         []string           `json:"zone" form:"zone" bson:"zone,omitempty"`
}
type AgroEcologicalZoneFilter struct {
	ActiveStatus []bool   `json:"activestatus,omitempty"`
	Status       []string `json:"status" form:"status" bson:"status,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	Searchbox    struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}
type RefAgroEcologicalZone struct {
	AgroEcologicalZone `bson:",inline"`
	Ref                struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
