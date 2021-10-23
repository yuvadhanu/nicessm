package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//CommonLand : "Holds single CommonLand data"
type CommonLand struct {
	ID primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`

	Status        string             `json:"status" bson:"status,omitempty"`
	ActiveStatus  bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Created       Created            `json:"createdOn" bson:"createdOn,omitempty"`
	GramPanchayat primitive.ObjectID `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	Village       primitive.ObjectID `json:"village"  bson:"village,omitempty"`
	Version       string             `json:"version"  bson:"version,omitempty"`
	Description   string             `json:"description"  bson:"description,omitempty"`
	State         primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	Block         primitive.ObjectID `json:"block"  bson:"block,omitempty"`
	District      primitive.ObjectID `json:"district"  bson:"district,omitempty"`
	Location      struct {
		Longitude string `json:"longitude" bson:"longitude,omitempty"`
		Latitude  string `json:"latitude" bson:"latitude,omitempty"`
	} `json:"location" bson:"location,omitempty"`
	Type         string `json:"type"  bson:"type,omitempty"`
	KhasraNumber string `json:"khasraNumber"  bson:"khasraNumber,omitempty"`
	Ownership    string `json:"ownership"  bson:"ownership,omitempty"`
	ParcelNumber string `json:"parcelNumber"  bson:"parcelNumber,omitempty"`
}

//RefCommonLand : "CommonLand with refrence data such as language..."
type RefCommonLand struct {
	CommonLand `bson:",inline"`
	Ref        struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//CommonLandFilter : "Used for constructing filter query"
type CommonLandFilter struct {
	ActiveStatus []bool   `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status       []string `json:"status" bson:"status,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	Regex        struct {
		Type         string `json:"type"  bson:"type"`
		KhasraNumber string `json:"khasraNumber"  bson:"khasraNumber"`
		Ownership    string `json:"ownership"  bson:"ownership"`
		ParcelNumber string `json:"parcelNumber"  bson:"parcelNumber"`
	} `json:"regex" bson:"regex"`
}
