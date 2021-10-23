package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CommoditySubVariety : ""
type CommoditySubVariety struct {
	ID               primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name             string             `json:"name,omitempty" bson:"name,omitempty"`
	Status           string             `json:"status,omitempty" bson:"status,omitempty"`
	CommodityVariety primitive.ObjectID `json:"commodityVariety" bson:"commodityVariety,omitempty"`
	Version          string             `json:"version,omitempty" bson:"version,omitempty"`
	Created          *CreatedV2         `json:"created,omitempty"  bson:"created,omitempty"`
}

type CommoditySubVarietyFilter struct {
	CommodityVariety []primitive.ObjectID `json:"commodityVariety" bson:"commodityVariety,omitempty"`
	Status           []string             `json:"status" form:"status" bson:"status,omitempty"`
	SortBy           string               `json:"sortBy"`
	SortOrder        int                  `json:"sortOrder"`
	SearchBox        struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchBox" bson:"searchBox"`
}

type RefCommoditySubVariety struct {
	CommoditySubVariety `bson:",inline"`
	Ref                 struct {
		CommodityVariety CommodityVariety `json:"commodityVariety" bson:"commodityVariety,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
