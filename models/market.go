package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Market : ""
type Market struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	ActiveStatus bool               `json:"activestatus" form:"activestatus" bson:"activestatus,omitempty"`
	Addressline1 string             `json:"addressline1" form:"addressline1" bson:"addressline1,omitempty"`
	Addressline2 string             `json:"addressline2" form:"addressline2" bson:"addressline2,omitempty"`
	Description  string             `json:"description" form:"description" bson:"description,omitempty"`
	District     primitive.ObjectID `json:"district" form:"district," bson:"district,omitempty"`
	Level        string             `json:"level " form:"level " bson:"level ,omitempty"`
	Location     struct {
		Latitude  float64 `json:"latitude" bson:"latitude"`
		Longitude float64 `json:"longitude" bson:"longitude"`
	} `json:"location" bson:"location"`
	Name          string             `json:"name" form:"name" bson:"name,omitempty"`
	State         primitive.ObjectID `json:"state" form:"state," bson:"state,omitempty"`
	Village       primitive.ObjectID `json:"village" form:"village," bson:"village,omitempty"`
	Block         primitive.ObjectID `json:"block" form:"block," bson:"block,omitempty"`
	GramPanchayat primitive.ObjectID `json:"gramPanchayat" form:"gramPanchayat," bson:"gramPanchayat,omitempty"`
	Status        string             `json:"status" form:"status" bson:"status,omitempty"`
	Created       *Created           `json:"created" form:"created" bson:"created,omitempty"`
	Version       int                `json:"version" form:"version" bson:"version,omitempty"`
}
type MarketFilter struct {
	ActiveStatus []bool          `json:"activestatus,omitempty"`
	Status       []string        `json:"status" form:"status" bson:"status,omitempty"`
	SortBy       string          `json:"sortBy"`
	SortOrder    int             `json:"sortOrder"`
	Address      AddressSearchV2 `json:"address,omitempty"`
	Regex        struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
}
type RefMarket struct {
	Market `bson:",inline"`
	Ref    struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
type AddressSearchV2 struct {
	StateCode          []primitive.ObjectID `json:"stateCode" bson:"stateCode,omitempty"`
	DistrictCode       []primitive.ObjectID `json:"districtCode" bson:"districtCode,omitempty"`
	BlockCode          []primitive.ObjectID `json:"blockCode" bson:"blockCode,omitempty"`
	GramPanchayathCode []primitive.ObjectID `json:"gramPanchayathCode" bson:"gramPanchayathCode,omitempty"`
	VillageCode        []primitive.ObjectID `json:"villageCode" bson:"villageCode,omitempty"`
	Country            []primitive.ObjectID `json:"country" bson:"country,omitempty"`
}
