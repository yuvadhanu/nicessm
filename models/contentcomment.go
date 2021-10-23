package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Commodity : ""
type ContentComment struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Comment      string             `json:"Comment,omitempty" bson:"Comment,omitempty"`
	ActiveStatus bool               `json:"activeStatus,omitempty" bson:"activeStatus,omitempty"`
	Status       string             `json:"status,omitempty" bson:"status,omitempty"`
	Content      primitive.ObjectID `json:"content,omitempty" bson:"content,omitempty"`
	DateCreated  *time.Time         `json:"dateCreated" form:"dateCreated" bson:"dateCreated,omitempty"`
	User         primitive.ObjectID `json:"User" bson:"User,omitempty"`
	Version      int                `json:"version" form:"version" bson:"version,omitempty"`
}
type ContentCommentFilter struct {
	Status    []string `json:"status" form:"status" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy"`
	SortOrder int      `json:"sortOrder"`
	SearchBox struct {
		Comment string `json:"Comment,omitempty" bson:"Comment,omitempty"`
	} `json:"searchBox" bson:"searchBox"`
}

type RefContentComment struct {
	ContentComment `bson:",inline"`
	Ref            struct {
		Content Content `json:"content,omitempty" bson:"content,omitempty"`
		User    User    `json:"User" bson:"User,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
