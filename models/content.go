package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Commodity : ""
type Content struct {
	ID              primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Class           string             `json:"class,omitempty" bson:"class,omitempty"`
	ActiveStatus    bool               `json:"activeStatus,omitempty" bson:"activeStatus,omitempty"`
	Author          primitive.ObjectID `json:"author" bson:"author,omitempty"`
	Content         string             `json:"content,omitempty" bson:"content,omitempty"`
	DateCreated     *time.Time         `json:"dateCreated" form:"dateCreated" bson:"dateCreated,omitempty"`
	IgnoredIndex    []string           `json:"ignoredIndex,omitempty" bson:"ignoredIndex,omitempty"`
	IndexingData    IndexingData       `json:"indexingData,omitempty" bson:"indexingData,omitempty"`
	KnowledgeDomain primitive.ObjectID `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
	LocationRank    int                `json:"locationRank,omitempty" bson:"locationRank,omitempty"`
	Source          string             `json:"source,omitempty" bson:"source,omitempty"`
	Status          string             `json:"status,omitempty" bson:"status,omitempty"`
	SubDomain       primitive.ObjectID `json:"subDomain" bson:"subDomain,omitempty"`
	SubTopic        primitive.ObjectID `json:"subTopic" bson:"subTopic,omitempty"`
	TimeApplicable  struct {
		Month              int    `json:"month" bson:"month,omitempty"`
		TimeApplicableType string `json:"timeApplicableType" bson:"timeApplicableType,omitempty"`
	} `json:"timeApplicable" bson:"timeApplicable,omitempty"`
	Topic           primitive.ObjectID `json:"Topic" bson:"Topic,omitempty"`
	Type            string             `json:"type,omitempty" bson:"type,omitempty"`
	Version         int                `json:"version" form:"version" bson:"version,omitempty"`
	RecordId        string             `json:"recordId,omitempty" bson:"recordId,omitempty"`
	DateReviewed    *time.Time         `json:"dateReviewed" form:"dateReviewed" bson:"dateReviewed,omitempty"`
	OriginalContent string             `json:"originalContent,omitempty" bson:"originalContent,omitempty"`
	ReviewedBy      string             `json:"reviewedBy,omitempty" bson:"reviewedBy,omitempty"`
	Organisation    primitive.ObjectID `json:"organisation" bson:"organisation,omitempty"`
	Project         primitive.ObjectID `json:"project" bson:"project,omitempty"`
	SmsType         string             `json:"smsType,omitempty" bson:"smsType,omitempty"`
	SmsContentType  string             `json:"smsContentType,omitempty" bson:"smsContentType,omitempty"`
}
type IndexingData struct {
	Classfication string               `json:"classfication,omitempty" bson:"classfication,omitempty"`
	State         []primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	District      []primitive.ObjectID `json:"district"  bson:"district,omitempty"`
	Block         []primitive.ObjectID `json:"block"  bson:"block,omitempty"`
	Season        []primitive.ObjectID `json:"season"  bson:"season,omitempty"`
}
type ContentFilter struct {
	Status    []string `json:"status" form:"status" bson:"status,omitempty"`
	Type      []string `json:"type,omitempty" bson:"type,omitempty"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
	SearchBox struct {
		Content string `json:"content,omitempty" bson:"content,omitempty"`
	} `json:"searchBox" bson:"searchBox"`
}

type RefContent struct {
	Content `bson:",inline"`
	Ref     struct {
		KnowledgeDomain KnowledgeDomain `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
		Organisation    Organisation    `json:"organisation" bson:"organisation,omitempty"`
		Project         Project         `json:"project" bson:"project,omitempty"`
		SubDomain       SubDomain       `json:"subDomain" bson:"subDomain,omitempty"`
		SubTopic        SubTopic        `json:"subTopic" bson:"subTopic,omitempty"`
		Topic           Topic           `json:"Topic" bson:"Topic,omitempty"`
		State           State           `json:"state"  bson:"state,omitempty"`
		District        District        `json:"district"  bson:"district,omitempty"`
		Block           Block           `json:"block"  bson:"block,omitempty"`
		Season          Cropseason      `json:"season"  bson:"season,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
