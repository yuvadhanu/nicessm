package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Topic : "Holds single Topic data"
type Topic struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	SubDomain    primitive.ObjectID `json:"subDomain" ," bson:"subDomain,omitempty"`
	Name         string             `json:"name" bson:"name,omitempty"`
	Status       string             `json:"status" bson:"status,omitempty"`
	ActiveStatus bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	DefaultValue struct {
		Classification string `json:"CLASSIFICATION"  bson:"CLASSIFICATION,omitempty"`
	} `json:"defaultValue"  bson:"defaultValue,omitempty"`
	IndexingParams     []string `json:"indexingParams"  bson:"indexingParams,omitempty"`
	TimeApplicableType string   `json:"timeApplicableType"  bson:"timeApplicableType,omitempty"`
	Created            Created  `json:"createdOn" bson:"createdOn,omitempty"`
	Version            string   `json:"version"  bson:"version,omitempty"`
	Description        string   `json:"description"  bson:"description,omitempty"`
}

//RefTopic : "Topic with refrence data such as language..."
type RefTopic struct {
	Topic `bson:",inline"`
	Ref   struct {
		KnowlegdeDomain KnowledgeDomain `json:"knowledgeDomain" ," bson:"knowledgeDomain,omitempty"`
		SubDomain       SubDomain       `json:"subDomain" ," bson:"subDomain,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//TopicFilter : "Used for constructing filter query"
type TopicFilter struct {
	SubDomain []primitive.ObjectID `json:"subDomain" ," bson:"subDomain,omitempty"`

	ActiveStatus []bool   `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status       []string `json:"status" bson:"status,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	Regex        struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
}
