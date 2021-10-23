package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//BlockCrop : "Holds single BlockCrop data"
type BlockCrop struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	State        primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	District     primitive.ObjectID `json:"district"  bson:"district,omitempty"`
	Block        primitive.ObjectID `json:"block"  bson:"block,omitempty"`
	Commodity    primitive.ObjectID `json:"commodity"  bson:"commodity,omitempty"` //filter [], lookup - commodity collection
	Name         string             `json:"name" bson:"name,omitempty"`
	Status       string             `json:"status" bson:"status,omitempty"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Created      Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Version      string             `json:"version"  bson:"version,omitempty"`

	NameInLocalLanguage string `json:"nameInLocalLanguage"  bson:"nameInLocalLanguage,omitempty"` // search text
}

//RefBlockCrop : ""
type RefBlockCrop struct {
	BlockCrop `bson:",inline"`
	Ref       struct {
		State     State     `json:"state,omitempty" bson:"state,omitempty"`
		District  District  `json:"district,omitempty" bson:"district,omitempty"`
		Block     Block     `json:"block,omitempty" bson:"block,omitempty"`
		Commodity Commodity `json:"commodity,omitempty" bson:"commodity,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//BlockCropFilter : "Used for constructing filter query"
type BlockCropFilter struct {
	State     []primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	Commodity []primitive.ObjectID `json:"commodity"  bson:"commodity,omitempty"` //filter [], lookup - commodity collection
	District  []primitive.ObjectID `json:"district"  bson:"district,omitempty"`
	Block     []primitive.ObjectID `json:"block"  bson:"block,omitempty"`

	ActiveStatus []bool   `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status       []string `json:"status" bson:"status,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	Regex        struct {
		Name                string `json:"name" bson:"name"`
		NameInLocalLanguage string `json:"nameInLocalLanguage"  bson:"nameInLocalLanguage,omitempty"`
	} `json:"regex" bson:"regex"`
}
