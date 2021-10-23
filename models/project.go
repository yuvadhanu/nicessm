package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Project : ""
type Project struct {
	ID                primitive.ObjectID   `json:"id" form:"id," bson:"_id,omitempty"`
	Name              string               `json:"name,omitempty"  bson:"name,omitempty"`
	ActiveStatus      bool                 `json:"activeStatus,omitempty"  bson:"activeStatus,omitempty"`
	Status            string               `json:"status,omitempty"  bson:"status,omitempty"`
	Budget            float64              `json:"budget,omitempty"  bson:"budget,omitempty"`
	NationalLevel     bool                 `json:"nationalLevel,omitempty"  bson:"nationalLevel,omitempty"`
	Mail              string               `json:"mail,omitempty"  bson:"mail,omitempty"`
	Organisation      primitive.ObjectID   `json:"organisation,omitempty" bson:"organisation,omitempty"`
	KnowledgeDomainID []primitive.ObjectID `json:"knowledgeDomainId,omitempty" bson:"knowledgeDomainId,omitempty"`
	StateID           []primitive.ObjectID `json:"stateId,omitempty" bson:"stateId,omitempty"`
	PartnerID         []primitive.ObjectID `json:"partnerId,omitempty" bson:"partnerId,omitempty"`
	Remarks           string               `json:"remarks,omitempty"  bson:"remarks,omitempty"`
	StartDate         *time.Time           `json:"startDate,omitempty"  bson:"startDate,omitempty"`
	EndDate           *time.Time           `json:"endDate,omitempty"  bson:"endDate,omitempty"`
	Version           float64              `json:"version,omitempty"  bson:"version,omitempty"`
	Created           *CreatedV2           `json:"created,omitempty"  bson:"created,omitempty"`
}

type ProjectFilter struct {
	Status        []string             `json:"status,omitempty" bson:"status,omitempty"`
	NationalLevel []bool               `json:"nationalLevel,omitempty"  bson:"nationalLevel,omitempty"`
	Organisation  []primitive.ObjectID `json:"organisation,omitempty"  bson:"organisation,omitempty"`
	BudgetRange   *struct {
		From float64 `json:"from"`
		To   float64 `json:"to"`
	} `json:"budgetRange"`
	StartDateRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"startDateRange"`
	EndDateRange *struct {
		From *time.Time `json:"from"`
		To   *time.Time `json:"to"`
	} `json:"endDateRange"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
	Regex     struct {
		Name string `json:"name" bson:"name"`
		Mail string `json:"mail,omitempty"  bson:"mail,omitempty"`
	} `json:"regex" bson:"regex"`
}

type RefProject struct {
	Project `bson:",inline"`
	Ref     struct {
		Organisation    Organisation         `json:"organisation,omitempty" bson:"organisation,omitempty"`
		States          []RefProjectState    `json:"states,omitempty" bson:"states,omitempty"`
		KnowledgeDomain []RefKnowledgeDomain `json:"knowledgeDomain,omitempty" bson:"knowledgeDomain,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
