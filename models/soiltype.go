package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//SoilType : ""
type SoilType struct {
	ID      primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty"  bson:"name,omitempty"`
	Version string             `json:"version,omitempty"  bson:"version,omitempty"`
	Status  string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type SoilTypeFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
	Regex     struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
}

type RefSoilType struct {
	SoilType `bson:",inline"`
	Ref      struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
