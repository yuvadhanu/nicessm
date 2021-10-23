package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CommodityVariety : ""
type CommodityVariety struct {
	ID        primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Status    string             `json:"status,omitempty" bson:"status,omitempty"`
	Commodity primitive.ObjectID `json:"commodity" bson:"commodity,omitempty"`
	Version   string             `json:"version,omitempty" bson:"version,omitempty"`
	Created   *CreatedV2         `json:"created,omitempty"  bson:"created,omitempty"`
}

type CommodityVarietyFilter struct {
	Commodity []primitive.ObjectID `json:"commodity" bson:"commodity,omitempty"`
	Status    []string             `json:"status" form:"status" bson:"status,omitempty"`
	SortBy    string               `json:"sortBy"`
	SortOrder int                  `json:"sortOrder"`
	SearchBox struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchBox" bson:"searchBox"`
}

type RefCommodityVariety struct {
	CommodityVariety `bson:",inline"`
	Ref              struct {
		Commodity           Commodity             `json:"commodity" bson:"commodity,omitempty"`
		CommoditySubVariety []CommoditySubVariety `json:"subVariety" bson:"subVariety,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
