package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//UserUserOrganisation : ""
type UserOrganisation struct {
	ID       primitive.ObjectID `json:"id" form:"id," bson:"id,omitempty"`
	UniqueID string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Status   string             `json:"status" bson:"status,omitempty"`
	Created  Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Updated  Updated            `json:"updated" form:"id," bson:"updated,omitempty"`
}

//RefUserUserOrganisation :""
type RefUserOrganisation struct {
	UserOrganisation `bson:",inline"`
	Ref              struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//UserUserOrganisationFilter : ""
type UserOrganisationFilter struct {
	Status    []string `json:"status"`
	UniqueID  []string `json:"uniqueId"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
	Regex     struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
}
