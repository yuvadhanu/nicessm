package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//LiveStockVaccination : "Holds single LiveStockVaccination data"
type LiveStockVaccination struct {
	ID                primitive.ObjectID   `json:"id" form:"id," bson:"_id,omitempty"`
	State             []primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	Name              string               `json:"name" bson:"name,omitempty"`
	Age               string               `json:"age" bson:"age,omitempty"`
	BoosterDose       string               `json:"boosterDose" bson:"boosterDose,omitempty"`
	BoosterTime       string               `json:"boosterTime" bson:"boosterTime,omitempty"`
	Dose              string               `json:"dose" bson:"dose,omitempty"`
	Immunity          string               `json:"immunity" bson:"immunity,omitempty"`
	Status            string               `json:"status" bson:"status,omitempty"`
	ActiveStatus      bool                 `json:"activeStatus" bson:"activeStatus,omitempty"`
	Booster           bool                 `json:"booster" bson:"booster,omitempty"`
	TimeOfVaccination bool                 `json:"timeOfVaccination" bson:"timeOfVaccination,omitempty"`
	Created           Created              `json:"createdOn" bson:"createdOn,omitempty"`
	Vaccine           primitive.ObjectID   `json:"vaccine"  bson:"vaccine,omitempty"`
	LiveStocks        []primitive.ObjectID `json:"liveStocks"  bson:"liveStocks"`
	Diseases          []primitive.ObjectID `json:"diseases"  bson:"diseases"`
	Version           string               `json:"version"  bson:"version,omitempty"`
	TimeTo            *time.Time           `json:"timeTo"  bson:"timeTo,omitempty"`
	TimeFrom          *time.Time           `json:"timeFrom"  bson:"timeFrom,omitempty"`
}

//RefLiveStockVaccination : ""
type RefLiveStockVaccination struct {
	LiveStockVaccination `bson:",inline"`
	Ref                  struct {
		Commodity []Commodity `json:"commodity,omitempty" bson:"commodity,omitempty"`
		State     []State     `json:"state"  bson:"state,omitempty"`
		Disease   []Disease   `json:"disease"  bson:"disease,omitempty"`
		Vaccine   Vaccine     `json:"vaccine"  bson:"vaccine,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//LiveStockVaccinationFilter : "Used for constructing filter query"
type LiveStockVaccinationFilter struct {
	State             []primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	Vaccine           primitive.ObjectID   `json:"vaccine"  bson:"vaccine,omitempty"`
	LiveStocks        []primitive.ObjectID `json:"liveStocks"  bson:"liveStocks"`
	Diseases          []primitive.ObjectID `json:"diseases"  bson:"diseases"`
	TimeOfVaccination []bool               `json:"timeOfVaccination" bson:"timeOfVaccination,omitempty"`
	ActiveStatus      []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Booster           []bool               `json:"booster" bson:"booster,omitempty"`
	Status            []string             `json:"status" bson:"status,omitempty"`
	SortBy            string               `json:"sortBy"`
	SortOrder         int                  `json:"sortOrder"`
	Regex             struct {
		Name        string `json:"name" bson:"name"`
		BoosterDose string `json:"boosterDose" bson:"boosterDose,omitempty"`
		BoosterTime string `json:"boosterTime" bson:"boosterTime,omitempty"`
		Dose        string `json:"dose" bson:"dose,omitempty"`
		Immunity    string `json:"immunity" bson:"immunity,omitempty"`
	} `json:"regex" bson:"regex"`
}
