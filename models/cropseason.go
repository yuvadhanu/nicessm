package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Cropseason : ""
type Cropseason struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Activestatus bool               `json:"activestatus" form:"activestatus" bson:"activestatus,omitempty"`
	Enddate      int                `json:"enddate" form:"enddate" bson:"enddate,omitempty"`
	Endmonth     int                `json:"endmonth" form:"endmonth" bson:"endmonth,omitempty"`
	Name         string             `json:"name" form:"name" bson:"name,omitempty"`
	Startdate    int                `json:"startdate" form:"startdate" bson:"startdate,omitempty"`
	Startmonth   int                `json:"startmonth" form:"startmonth" bson:"startmonth,omitempty"`
	State        primitive.ObjectID `json:"state" form:"state," bson:"state,omitempty"`
	Status       string             `json:"status" form:"status" bson:"status,omitempty"`
	Created      *Created           `json:"created" form:"created" bson:"created,omitempty"`
	Version      int                `json:"version" form:"version" bson:"version,omitempty"`
	YearPlusOne  bool               `json:"yearPlusOne" form:"yearPlusOne" bson:"yearPlusOne,omitempty"`
}
type CropseasonFilter struct {
	ActiveStatus []bool               `json:"activestatus,omitempty"`
	Status       []string             `json:"status" form:"status" bson:"status,omitempty"`
	SortBy       string               `json:"sortBy"`
	SortOrder    int                  `json:"sortOrder"`
	State        []primitive.ObjectID `json:"state" bson:"state,"`
	Searchbox    struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}
type RefCropseason struct {
	Cropseason `bson:",inline"`
	Ref        struct {
		State State `json:"state" form:"state," bson:"state,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
