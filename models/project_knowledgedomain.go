package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ProjectKnowledgeDomain : ""
type ProjectKnowledgeDomain struct {
	ID              primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	KnowledgeDomain primitive.ObjectID `json:"knowledgeDomain" form:"knowledgeDomain," bson:"knowledgeDomain,omitempty"`
	Project         primitive.ObjectID `json:"project" form:"project," bson:"project,omitempty"`
	Status          string             `json:"status,omitempty" bson:"status,omitempty"`
	Created         *CreatedV2         `json:"created,omitempty"  bson:"created,omitempty"`
}

type ProjectKnowledgeDomainFilter struct {
	Status          []string             `json:"status,omitempty" bson:"status,omitempty"`
	Project         []primitive.ObjectID `json:"project,omitempty" bson:"project,omitempty"`
	KnowledgeDomain []primitive.ObjectID `json:"knowledgeDomain,omitempty" bson:"knowledgeDomain,omitempty"`
	SortBy          string               `json:"sortBy"`
	SortOrder       int                  `json:"sortOrder"`
}

type RefProjectKnowledgeDomain struct {
	ProjectKnowledgeDomain `bson:",inline"`
	Ref                    struct {
		Project         Project         `json:"project" bson:"project,omitempty"`
		KnowledgeDomain KnowledgeDomain `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
