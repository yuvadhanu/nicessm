package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//SubTopic : "Holds single SubTopic data"
type SubTopic struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Topic        primitive.ObjectID `json:"topic" ," bson:"topic,omitempty"`
	Name         string             `json:"name" bson:"name,omitempty"`
	Status       string             `json:"status" bson:"status,omitempty"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Created      Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Version      string             `json:"version"  bson:"version,omitempty"`
	Description  string             `json:"description"  bson:"description,omitempty"`
}

//RefSubTopic : "SubTopic with refrence data such as language..."
type RefSubTopic struct {
	SubTopic `bson:",inline"`
	Ref      struct {
		KnowlegdeDomain KnowledgeDomain `json:"knowledgeDomain" ," bson:"knowledgeDomain,omitempty"`
		SubDomain       SubDomain       `json:"subDomain" ," bson:"subDomain,omitempty"`
		Topic           Topic           `json:"topic" ," bson:"topic,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//SubTopicFilter : "Used for constructing filter query"
type SubTopicFilter struct {
	Topic        []primitive.ObjectID `json:"topic" ," bson:"topic,omitempty"`
	ActiveStatus []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status       []string             `json:"status" bson:"status,omitempty"`
	SortBy       string               `json:"sortBy"`
	SortOrder    int                  `json:"sortOrder"`
	Regex        struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
}
