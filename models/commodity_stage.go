package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CommodityStage : ""
type CommodityStage struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	Status         string             `json:"status,omitempty" bson:"status,omitempty"`
	Category       primitive.ObjectID `json:"category" bson:"category,omitempty"`
	Commodity      primitive.ObjectID `json:"commodity" bson:"commodity,omitempty"`
	ActiveStatus   bool               `json:"activeStatus,omitempty" bson:"activeStatus,omitempty"`
	SequenceNumber float64            `json:"sequenceNumber,omitempty" bson:"sequenceNumber,omitempty"`
	Version        string             `json:"version,omitempty" bson:"version,omitempty"`
	Created        *CreatedV2         `json:"created,omitempty"  bson:"created,omitempty"`
}

type CommodityStageFilter struct {
	Category  []primitive.ObjectID `json:"category" bson:"category,omitempty"`
	Commodity      []primitive.ObjectID `json:"commodity" bson:"commodity,omitempty"`
	Status    []string             `json:"status" form:"status" bson:"status,omitempty"`
	SortBy    string               `json:"sortBy"`
	SortOrder int                  `json:"sortOrder"`
	SearchBox struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchBox" bson:"searchBox"`
}

type RefCommodityStage struct {
	CommodityStage `bson:",inline"`
	Ref            struct {
		Category  CommodityCategory `json:"category" bson:"category,omitempty"`
		Commodity Commodity         `json:"commodity" bson:"commodity,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
