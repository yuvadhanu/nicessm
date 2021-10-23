package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Commodity : ""
type ContentTranslation struct {
	ID                primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Content           primitive.ObjectID `json:"content,omitempty" bson:"content,omitempty"`
	DateCreated       *time.Time         `json:"dateCreated" form:"dateCreated" bson:"dateCreated,omitempty"`
	Language          string             `json:"language,omitempty" bson:"language,omitempty"`
	ActiveStatus      bool               `json:"activeStatus,omitempty" bson:"activeStatus,omitempty"`
	Status            string             `json:"status,omitempty" bson:"status,omitempty"`
	Translator        primitive.ObjectID `json:"translator" bson:"translator,omitempty"`
	TranslatedContent string             `json:"translatedContent,omitempty" bson:"translatedContent,omitempty"`
	Version           int                `json:"version" form:"version" bson:"version,omitempty"`
	DateReviewed      *time.Time         `json:"dateReviewed" form:"dateReviewed" bson:"dateReviewed,omitempty"`
	ReviewedBy        primitive.ObjectID `json:"reviewedBy" bson:"reviewedBy,omitempty"`
}
type ContentTranslationFilter struct {
	Status    []string `json:"status" form:"status" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
	SearchBox struct {
		//	Content           primitive.ObjectID `json:"content,omitempty" bson:"content,omitempty"`
		TranslatedContent string `json:"translatedContent,omitempty" bson:"translatedContent,omitempty"`
	} `json:"searchBox" bson:"searchBox"`
}

type RefContentTranslation struct {
	ContentTranslation `bson:",inline"`
	Ref                struct {
		Content    Content `json:"content,omitempty" bson:"content,omitempty"`
		Translator User    `json:"translator" bson:"translator,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
