package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//State : "Holds single state data"
type State struct {
	ID primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`

	Name         string               `json:"name" bson:"name,omitempty"`
	Status       string               `json:"status" bson:"status,omitempty"`
	ActiveStatus bool                 `json:"activeStatus" bson:"activeStatus,omitempty"`
	Created      Created              `json:"createdOn" bson:"createdOn,omitempty"`
	Languages    []primitive.ObjectID `json:"languages"  bson:"languages,omitempty"`
	Version      string               `json:"version"  bson:"version,omitempty"`
}

//RefState : "State with refrence data such as language..."
type RefState struct {
	State `bson:",inline"`
	Ref   struct {
		Languages []Languages `json:"languages,omitempty" bson:"languages,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//StateFilter : "Used for constructing filter query"
type StateFilter struct {
	ActiveStatus []bool   `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status       []string `json:"status" bson:"status,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	Regex        struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
}
