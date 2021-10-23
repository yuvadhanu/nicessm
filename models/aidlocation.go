package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Language : ""
type Aidlocation struct {
	ID          primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Address     string             `json:"address" form:"address" bson:"address,omitempty"`
	Description string             `json:"description" form:"description" bson:"description,omitempty"`
	Location    struct {
		Latitude  float64 `json:"latitude" bson:"latitude"`
		Longitude float64 `json:"longitude" bson:"longitude"`
	} `json:"location" bson:"location"`
	Name        string               `json:"name" form:"name" bson:"name,omitempty"`
	Phoneno     string               `json:"phoneno" form:"phoneno" bson:"phoneno,omitempty"`
	Status      string               `json:"status" form:"status" bson:"status,omitempty"`
	Created     *Created             `json:"created" form:"created" bson:"created,omitempty"`
	Version     int                  `json:"version" form:"version" bson:"version,omitempty"`
	AidCategory []primitive.ObjectID `json:"aidCategory" form:"aidCategory," bson:"aidCategory,omitempty"`
}
type AidlocationFilter struct {
	Status    []string `json:"status" form:"status" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
	Searchbox struct {
		Name        string `json:"name" bson:"name"`
		Description string `json:"description" form:"description" bson:"description,omitempty"`
	} `json:"searchbox" bson:"searchbox"`
}
type RefAidlocation struct {
	Aidlocation `bson:",inline"`
	Ref         struct {
		AidCategory []AidCategory `json:"aidCategory" form:"aidCategory," bson:"aidCategory,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
