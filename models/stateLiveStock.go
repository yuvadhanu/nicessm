package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//StateLiveStock : "Holds single StateLiveStock data"
type StateLiveStock struct {
	ID                  primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	State               primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	Name                string             `json:"name" bson:"name,omitempty"`
	Status              string             `json:"status" bson:"status,omitempty"`
	ActiveStatus        bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Created             Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Version             string             `json:"version"  bson:"version,omitempty"`
	Commodity           primitive.ObjectID `json:"commodity"  bson:"commodity,omitempty"`                     //filter [], lookup - commodity collection
	NameInLocalLanguage string             `json:"nameInLocalLanguage"  bson:"nameInLocalLanguage,omitempty"` // search text
}

//RefStateLiveStock : ""
type RefStateLiveStock struct {
	StateLiveStock `bson:",inline"`
	Ref            struct {
		State     State       `json:"state,omitempty" bson:"state,omitempty"`
		Commodity []Commodity `json:"commodity,omitempty" bson:"commodity,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//StateLiveStockFilter : "Used for constructing filter query"
type StateLiveStockFilter struct {
	State        []primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	Commodity    []primitive.ObjectID `json:"commodity"  bson:"commodity,omitempty"` //filter [], lookup - commodity collection
	ActiveStatus []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status       []string             `json:"status" bson:"status,omitempty"`
	SortBy       string               `json:"sortBy"`
	SortOrder    int                  `json:"sortOrder"`
	Regex        struct {
		Name                string `json:"name" bson:"name"`
		NameInLocalLanguage string `json:"nameInLocalLanguage"  bson:"nameInLocalLanguage,omitempty"`
	} `json:"regex" bson:"regex"`
}
