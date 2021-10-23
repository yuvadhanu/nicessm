package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//KnowlegdeDomain : "Holds single KnowlegdeDomain data"
type KnowledgeDomain struct {
	ID primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`

	Name         string  `json:"name" bson:"name,omitempty"`
	Status       string  `json:"status" bson:"status,omitempty"`
	ActiveStatus bool    `json:"activeStatus" bson:"activeStatus,omitempty"`
	Created      Created `json:"createdOn" bson:"createdOn,omitempty"`

	Version     string `json:"version"  bson:"version,omitempty"`
	Description string `json:"description"  bson:"description,omitempty"`
}

//RefKnowlegdeDomain : "KnowlegdeDomain with refrence data such as language..."
type RefKnowledgeDomain struct {
	KnowledgeDomain `bson:",inline"`
	Ref             struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//KnowlegdeDomainFilter : "Used for constructing filter query"
type KnowledgeDomainFilter struct {
	ActiveStatus []bool   `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status       []string `json:"status" bson:"status,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	Regex        struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
}
