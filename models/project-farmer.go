package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ProjectFarmer : ""
type ProjectFarmer struct {
	ID      primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Farmer  primitive.ObjectID `json:"farmer" bson:"farmer,omitempty"`
	Project primitive.ObjectID `json:"project" form:"project," bson:"project,omitempty"`
	Status  string             `json:"status,omitempty" bson:"status,omitempty"`
	Created *CreatedV2         `json:"created,omitempty"  bson:"created,omitempty"`
}

type ProjectFarmerFilter struct {
	Status    []string             `json:"status,omitempty" bson:"status,omitempty"`
	Project   []primitive.ObjectID `json:"project,omitempty" bson:"project,omitempty"`
	Farmer    []primitive.ObjectID `json:"farmer,omitempty" bson:"farmer,omitempty"`
	SortBy    string               `json:"sortBy"`
	SortOrder int                  `json:"sortOrder"`
}

type RefProjectFarmer struct {
	ProjectFarmer `bson:",inline"`
	Ref           struct {
		Project Project `json:"project" bson:"project,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
