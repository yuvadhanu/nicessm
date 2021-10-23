package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ProjectState : ""
type ProjectState struct {
	ID      primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	State   primitive.ObjectID `json:"state" form:"state," bson:"state,omitempty"`
	Project primitive.ObjectID `json:"project" form:"project," bson:"project,omitempty"`
	Status  string             `json:"status,omitempty" bson:"status,omitempty"`
	Created *CreatedV2         `json:"created,omitempty"  bson:"created,omitempty"`
}

type ProjectStateFilter struct {
	Status    []string             `json:"status,omitempty" bson:"status,omitempty"`
	Project   []primitive.ObjectID `json:"project,omitempty" bson:"project,omitempty"`
	State     []primitive.ObjectID `json:"state,omitempty" bson:"state,omitempty"`
	SortBy    string               `json:"sortBy"`
	SortOrder int                  `json:"sortOrder"`
}

type RefProjectState struct {
	ProjectState `bson:",inline"`
	Ref          struct {
		Project Project `json:"project" bson:"project,omitempty"`
		State   State   `json:"state" bson:"state,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
