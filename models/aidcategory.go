package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//AidCategory : ""
type AidCategory struct {
	ID      primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty"  bson:"name,omitempty"`
	Version int                `json:"version,omitempty"  bson:"version,omitempty"`
	Status  string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type AidCategoryFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
	SearchBox struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefAidCategory struct {
	AidCategory `bson:",inline"`
	Ref         struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
