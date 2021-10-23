package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//SoilType : ""
type Asset struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty"  bson:"name,omitempty"`
	ActiveStatus bool               `json:"activestatus"  bson:"activestatus,omitempty"`
	Version      string             `json:"version,omitempty"  bson:"version,omitempty"`
	Status       string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created      *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type AssetFilter struct {
	Status       []string `json:"status,omitempty" bson:"status,omitempty"`
	ActiveStatus []bool   `json:"activestatus"  bson:"activestatus,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	Regex        struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
}

type RefAsset struct {
	Asset `bson:",inline"`
	Ref   struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
