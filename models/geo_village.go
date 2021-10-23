package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Village : "Holds single state data"
type Village struct {
	ID            primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	GramPanchayat primitive.ObjectID `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	Name          string             `json:"name" bson:"name,omitempty"`
	Status        string             `json:"status" bson:"status,omitempty"`
	ActiveStatus  bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Created       CreatedV2          `json:"createdOn" bson:"createdOn,omitempty"`
	Location      struct {
		Longitude float64 `json:"longitude" bson:"longitude,omitempty"`
		Latitude  float64 `json:"latitude" bson:"latitude,omitempty"`
	} `json:"location" bson:"location,omitempty"`
	Population      string `json:"population"  bson:"population,omitempty"`
	VillageHead     string `json:"villageHead"  bson:"villageHead,omitempty"`
	CommiteeDetails string `json:"commiteeDetails"  bson:"commiteeDetails,omitempty"`
	School          string `json:"school"  bson:"school,omitempty"`
	FieldAgent      string `json:"fieldAgent"  bson:"fieldAgent,omitempty"`
	Version         string `json:"version"  bson:"version,omitempty"`
}

//RefVillage : "Village with refrence data such as language..."
type RefVillage struct {
	Village `bson:",inline"`
	Ref     struct {
		State         State         `json:"state,omitempty" bson:"state,omitempty"`
		District      District      `json:"district,omitempty" bson:"district,omitempty"`
		Block         Block         `json:"block,omitempty" bson:"block,omitempty"`
		GramPanchayat GramPanchayat `json:"gramPanchayat,omitempty" bson:"gramPanchayat,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//VillageFilter : "Used for constructing filter query"
type VillageFilter struct {
	GramPanchayat []primitive.ObjectID `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	ActiveStatus  []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status        []string             `json:"status" bson:"status,omitempty"`
	SortBy        string               `json:"sortBy"`
	SortOrder     int                  `json:"sortOrder"`
	Regex         struct {
		Name string `json:"name" bson:"name"`
		Type string `json:"type"  bson:"type"`
	} `json:"regex" bson:"regex"`
}
