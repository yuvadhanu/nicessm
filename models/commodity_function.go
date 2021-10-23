package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CommodityFunction : ""
type CommodityFunction struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name,omitempty"`
	Status       string             `json:"status,omitempty" bson:"status,omitempty"`
	Category     primitive.ObjectID `json:"category" bson:"category,omitempty"`
	ActiveStatus bool               `json:"activeStatus,omitempty" bson:"activeStatus,omitempty"`
	Version      string             `json:"version,omitempty" bson:"version,omitempty"`
	Created      *CreatedV2         `json:"created,omitempty"  bson:"created,omitempty"`
}

type CommodityFunctionFilter struct {
	Category  []primitive.ObjectID `json:"category" bson:"category,omitempty"`
	Status    []string             `json:"status" form:"status" bson:"status,omitempty"`
	SortBy    string               `json:"sortBy"`
	SortOrder int                  `json:"sortOrder"`
	SearchBox struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchBox" bson:"searchBox"`
}

type RefCommodityFunction struct {
	CommodityFunction `bson:",inline"`
	Ref               struct {
		Category CommodityCategory `json:"category" bson:"category,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
