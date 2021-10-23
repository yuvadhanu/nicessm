package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//AidCategory : ""
type DistrictWeatherData struct {
	ID                      primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	Name                    string             `json:"name" bson:"name"`
	ActiveStatus            bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	District                string             `json:"district,omitempty"  bson:"district,omitempty"`
	State                   string             `json:"state" bson:"state,omitempty"`
	MaxRelativeHumidityFrom float64            `json:"maxRelativeHumidityFrom,omitempty"  bson:"maxRelativeHumidityFrom,omitempty"`
	MaxRelativeHumidityTo   float64            `json:"maxRelativeHumidityTo,omitempty"  bson:"maxRelativeHumidityTo,omitempty"`
	MaxTemperatureFrom      float64            `json:"maxTemperatureFrom,omitempty"  bson:"maxTemperatureFrom,omitempty"`
	MaxTemperatureTo        float64            `json:"maxTemperatureTo,omitempty"  bson:"maxTemperatureTo,omitempty"`
	MinRelativeHumidityFrom float64            `json:"minRelativeHumidityFrom,omitempty"  bson:"minRelativeHumidityFrom,omitempty"`
	MinRelativeHumidityTo   float64            `json:"minRelativeHumidityTo,omitempty"  bson:"minRelativeHumidityTo,omitempty"`
	MinTemperatureFrom      float64            `json:"minTemperatureFrom,omitempty"  bson:"minTemperatureFrom,omitempty"`
	MinTemperatureTo        float64            `json:"minTemperatureTo,omitempty"  bson:"minTemperatureTo,omitempty"`
	RainfallFrom            float64            `json:"rainfallFrom,omitempty"  bson:"rainfallFrom,omitempty"`
	RainfallTo              float64            `json:"rainfallTo,omitempty"  bson:"rainfallTo,omitempty"`
	Version                 int                `json:"version,omitempty"  bson:"version,omitempty"`
	WindDirectionFrom       float64            `json:"windDirectionFrom,omitempty"  bson:"windDirectionFrom,omitempty"`
	WindDirectionTo         float64            `json:"windDirectionTo,omitempty"  bson:"windDirectionTo,omitempty"`
	WindSpeedFrom           float64            `json:"windSpeedFrom,omitempty"  bson:"windSpeedFrom,omitempty"`
	WindSpeedTo             float64            `json:"windSpeedTo,omitempty"  bson:"windSpeedTo,omitempty"`
	Status                  string             `json:"status,omitempty"  bson:"status,omitempty"`
	Created                 *Created           `json:"created,omitempty"  bson:"created,omitempty"`
}

type DistrictWeatherDataFilter struct {
	Status       []string `json:"status,omitempty" bson:"status,omitempty"`
	ActiveStatus []bool   `json:"activeStatus" bson:"activeStatus,omitempty"`
	District     []string `json:"district,omitempty"  bson:"district,omitempty"`
	State        []string `json:"state" bson:"state,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	SearchBox    struct {
		Name string `json:"name" bson:"name"`
	} `json:"searchbox" bson:"searchbox"`
}

type RefDistrictWeatherData struct {
	DistrictWeatherData `bson:",inline"`
	Ref                 struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
