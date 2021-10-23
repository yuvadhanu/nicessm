package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Organisation : ""
type Organisation struct {
	ID primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	//UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name         string  `json:"name" bson:"name,omitempty"`
	Status       string  `json:"status" bson:"status,omitempty"`
	ActiveStatus bool    `json:"activeStatus" bson:"activeStatus,omitempty"`
	Created      Created `json:"createdOn" bson:"createdOn,omitempty"`
	//Updated     Updated `json:"updated"  bson:"updated,omitempty"` version
	Description string   `json:"description" bson:"description,omitempty"`
	Tags        []string `json:"tags"  bson:"tags,omitempty"`
	Version     string   `json:"version"  bson:"version,omitempty"`
}

//RefOrganisation :""
type RefOrganisation struct {
	Organisation `bson:",inline"`
	Ref          struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//OrganisationFilter : ""

type OrganisationFilter struct {
	ActiveStatus []bool   `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status       []string `json:"status" bson:"status,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	Regex        struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
}
