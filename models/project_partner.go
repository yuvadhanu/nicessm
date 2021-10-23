package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ProjectPartner : ""
type ProjectPartner struct {
	ID      primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Partner primitive.ObjectID `json:"partner" bson:"partner,omitempty"`
	Project primitive.ObjectID `json:"project" form:"project," bson:"project,omitempty"`
	Status  string             `json:"status,omitempty" bson:"status,omitempty"`
	Created *CreatedV2         `json:"created,omitempty"  bson:"created,omitempty"`
}

type ProjectPartnerFilter struct {
	Status    []string             `json:"status,omitempty" bson:"status,omitempty"`
	Project   []primitive.ObjectID `json:"project,omitempty" bson:"project,omitempty"`
	Partner   []primitive.ObjectID `json:"partner,omitempty" bson:"partner,omitempty"`
	SortBy    string               `json:"sortBy"`
	SortOrder int                  `json:"sortOrder"`
}

type RefProjectPartner struct {
	ProjectPartner `bson:",inline"`
	Ref            struct {
		Project Project `json:"project" bson:"project,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
