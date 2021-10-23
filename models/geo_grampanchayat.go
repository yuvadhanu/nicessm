package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//GramPanjayat : "Holds single state data"
type GramPanchayat struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Block        primitive.ObjectID `json:"block"  bson:"block,omitempty"`
	Name         string             `json:"name" bson:"name,omitempty"`
	Status       string             `json:"status" bson:"status,omitempty"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Created      CreatedV2          `json:"createdOn" bson:"createdOn,omitempty"`
	//Languages    []primitive.ObjectID `json:"languages"  bson:"languages,omitempty"`
	Version string `json:"version"  bson:"version,omitempty"`
}

//RefGramPanjayat : "Village with refrence data such as language..."
type RefGramPanchayat struct {
	GramPanchayat `bson:",inline"`
	Ref           struct {
		State    State    `json:"state,omitempty" bson:"state,omitempty"`
		District District `json:"district,omitempty" bson:"district,omitempty"`
		Block    Block    `json:"block,omitempty" bson:"block,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//GramPanjayatFilter : "Used for constructing filter query"
type GramPanchayatFilter struct {
	Block        []primitive.ObjectID `json:"block"  bson:"block,omitempty"`
	ActiveStatus []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status       []string             `json:"status" bson:"status,omitempty"`
	SortBy       string               `json:"sortBy"`
	SortOrder    int                  `json:"sortOrder"`
	Regex        struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
}
