package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Block : "Holds single state data"
type Block struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	District     primitive.ObjectID `json:"district"  bson:"district,omitempty"`
	Name         string             `json:"name" bson:"name,omitempty"`
	Status       string             `json:"status" bson:"status,omitempty"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Created      CreatedV2          `json:"createdOn" bson:"createdOn,omitempty"`
	//Languages    []primitive.ObjectID `json:"languages"  bson:"languages,omitempty"`
	Version string `json:"version"  bson:"version,omitempty"`
}

//RefBlock : "Village with refrence data such as language..."
type RefBlock struct {
	Block `bson:",inline"`
	Ref   struct {
		State    State    `json:"state,omitempty" bson:"state,omitempty"`
		District District `json:"district,omitempty" bson:"district,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//BlockFilter : "Used for constructing filter query"
type BlockFilter struct {
	District []primitive.ObjectID `json:"district"  bson:"district,omitempty"`

	ActiveStatus []bool   `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status       []string `json:"status" bson:"status,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	Regex        struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
}
