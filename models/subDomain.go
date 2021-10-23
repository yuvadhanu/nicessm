package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//SubDomain : "Holds single SubDomain data"
type SubDomain struct {
	ID              primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	KnowledgeDomain primitive.ObjectID `json:"knowledgeDomain" ," bson:"knowledgeDomain,omitempty"`
	Name            string             `json:"name" bson:"name,omitempty"`
	Status          string             `json:"status" bson:"status,omitempty"`
	ActiveStatus    bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Created         Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Version         string             `json:"version"  bson:"version,omitempty"`
	Description     string             `json:"description"  bson:"description,omitempty"`
}

//RefSubDomain : "SubDomain with refrence data such as language..."
type RefSubDomain struct {
	SubDomain `bson:",inline"`
	Ref       struct {
		KnowledgeDomain KnowledgeDomain `json:"knowledgeDomain" ," bson:"knowledgeDomain,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//SubDomainFilter : "Used for constructing filter query"
type SubDomainFilter struct {
	KnowledgeDomain []primitive.ObjectID `json:"knowledgeDomain" ," bson:"knowledgeDomain,omitempty"`
	ActiveStatus    []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status          []string             `json:"status" bson:"status,omitempty"`
	SortBy          string               `json:"sortBy"`
	SortOrder       int                  `json:"sortOrder"`
	Regex           struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
}
