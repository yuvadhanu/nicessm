package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ProjectUser : ""
type ProjectUser struct {
	ID      primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	User    primitive.ObjectID `json:"user" form:"user," bson:"user,omitempty"`
	Project primitive.ObjectID `json:"project" form:"project," bson:"project,omitempty"`
	Status  string             `json:"status,omitempty" bson:"status,omitempty"`
	Created *CreatedV2         `json:"created,omitempty"  bson:"created,omitempty"`
}

type ProjectUserFilter struct {
	Status    []string             `json:"status,omitempty" bson:"status,omitempty"`
	Project   []primitive.ObjectID `json:"project,omitempty" bson:"project,omitempty"`
	User      []primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	SortBy    string               `json:"sortBy"`
	SortOrder int                  `json:"sortOrder"`
}

type RefProjectUser struct {
	ProjectUser `bson:",inline"`
	Ref         struct {
		Project Project `json:"project" bson:"project,omitempty"`
		User    User    `json:"user" bson:"user,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
