package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//AidCategory : ""
type Disease struct {
	ID             primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus   bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Name           string             `json:"name,omitempty"  bson:"name,omitempty"`
	Classification string             `json:"classification" bson:"classification,omitempty"`
	Version        int                `json:"version,omitempty"  bson:"version,omitempty"`
	Status         string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created        *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type DiseaseFilter struct {
	OmitIDs      []primitive.ObjectID `json:"omitIds,omitempty"`
	Status       []string             `json:"status,omitempty" bson:"status,omitempty"`
	ActiveStatus []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	SortBy       string               `json:"sortBy"`
	SortOrder    int                  `json:"sortOrder"`
	SearchBox    struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefDisease struct {
	Disease `bson:",inline"`
	Ref     struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
