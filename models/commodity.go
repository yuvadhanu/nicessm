package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Commodity : ""
type Commodity struct {
	ID             primitive.ObjectID   `json:"id" form:"id," bson:"_id,omitempty"`
	Name           string               `json:"name,omitempty" bson:"name,omitempty"`
	CommonName     string               `json:"commonName,omitempty" bson:"commonName,omitempty"`
	ScientificName string               `json:"scientificName,omitempty" bson:"scientificName,omitempty"`
	Status         string               `json:"status,omitempty" bson:"status,omitempty"`
	Category       primitive.ObjectID   `json:"category" bson:"category,omitempty"`
	Function       primitive.ObjectID   `json:"function" bson:"function,omitempty"`
	Insects        []primitive.ObjectID `json:"insects" bson:"insects,omitempty"`
	Diseases       []primitive.ObjectID `json:"diseases" bson:"diseases,omitempty"`
	ActiveStatus   bool                 `json:"activeStatus,omitempty" bson:"activeStatus,omitempty"`
	Version        string               `json:"version,omitempty" bson:"version,omitempty"`
	Created        *CreatedV2           `json:"created,omitempty"  bson:"created,omitempty"`
}

type CommodityFilter struct {
	Function      []primitive.ObjectID `json:"function" bson:"function,omitempty"`
	Category      []primitive.ObjectID `json:"category" bson:"category,omitempty"`
	Status        []string             `json:"status" form:"status" bson:"status,omitempty"`
	Classfication []string             `json:"classfication" form:"classfication" bson:"classfication,omitempty"`
	SortBy        string               `json:"sortBy"`
	SortOrder     int                  `json:"sortOrder"`
	SearchBox     struct {
		CommonName     string `json:"commonName" bson:"commonName"`
		ScientificName string `json:"scientificName" bson:"scientificName"`
	} `json:"searchBox" bson:"searchBox"`
}

type RefCommodity struct {
	Commodity `bson:",inline"`
	Ref       struct {
		Category CommodityCategory `json:"category" bson:"category,omitempty"`
		Function CommodityFunction `json:"function" bson:"function,omitempty"`
		Insects  []Insect          `json:"insects" bson:"insects,omitempty"`
		Diseases []Disease         `json:"diseases" bson:"diseases,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
