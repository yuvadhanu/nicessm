package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Farmer : "Holds single Farmer data"
type Farmer struct {
	ID                           primitive.ObjectID  `json:"id" form:"id," bson:"_id,omitempty"`
	AlternateNumber              string              `json:"alternateNumber" bson:"alternateNumber,omitempty"`
	Education                    string              `json:"education" bson:"education,omitempty"`
	FatherName                   string              `json:"fatherName" bson:"fatherName,omitempty"`
	CurrentLiveStocks            []CurrentLiveStocks `json:"currentLiveStocks" bson:"currentLiveStocks,omitempty"`
	CurrentCrops                 []CurrentCrops      `json:"currentCrops" bson:"currentCrops,omitempty"`
	CultivationPractice          string              `json:"cultivationPractice" bson:"cultivationPractice,omitempty"`
	Name                         string              `json:"name" bson:"name,omitempty"`
	Gender                       string              `json:"gender" bson:"gender,omitempty"`
	DateOfBirth                  *time.Time          `json:"dateOfBirth" bson:"dateOfBirth,omitempty"`
	Status                       string              `json:"status" bson:"status,omitempty"`
	District                     primitive.ObjectID  `json:"district"  bson:"district,omitempty"`
	GramPanchayat                primitive.ObjectID  `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	IsMemberInvolvedInCbo        bool                `json:"isMemberInvolvedInCbo" bson:"isMemberInvolvedInCbo,omitempty"`
	ActiveStatus                 bool                `json:"activeStatus" bson:"activeStatus,omitempty"`
	IsSMS                        bool                `json:"isSMS" bson:"isSMS,omitempty"`
	FeminineMobile               bool                `json:"feminineMobile" bson:"feminineMobile,omitempty"`
	CreditAvailed                bool                `json:"creditAvailed" bson:"creditAvailed,omitempty"`
	Created                      Created             `json:"createdOn" bson:"createdOn,omitempty"`
	Block                        primitive.ObjectID  `json:"block"  bson:"block,omitempty"`
	Version                      string              `json:"version"  bson:"version,omitempty"`
	IsVoiceSMS                   bool                `json:"isVoiceSMS" bson:"isVoiceSMS,omitempty"`
	KitchenGarden                bool                `json:"kitchenGarden" bson:"kitchenGarden,omitempty"`
	LeasedInIrrigated            float64             `json:"leasedInIrrigated"  bson:"leasedInIrrigated,omitempty"`
	LeasedInRainfed              float64             `json:"leasedInRainfed"  bson:"leasedInRainfed,omitempty"`
	LeasedOutIrrigated           float64             `json:"leasedOutIrrigated"  bson:"leasedOutIrrigated,omitempty"`
	LeasedOutRainfed             float64             `json:"leasedOutRainfed"  bson:"leasedOutRainfed,omitempty"`
	MembershipInMgnrega          bool                `json:"membershipInMgnrega" bson:"membershipInMgnrega,omitempty"`
	MobileNumber                 string              `json:"mobileNumber"  bson:"mobileNumber,omitempty"`
	OwnedIrrigated               float64             `json:"ownedIrrigated"  bson:"ownedIrrigated,omitempty"`
	OwnedRainfed                 float64             `json:"ownedRainfed"  bson:"ownedRainfed,omitempty"`
	SoilType                     primitive.ObjectID  `json:"soilType"  bson:"soilType,omitempty"`
	SpouseName                   string              `json:"spouseName"  bson:"spouseName,omitempty"`
	State                        primitive.ObjectID  `json:"state"  bson:"state,omitempty"`
	TotalLand                    float64             `json:"totalLand"  bson:"totalLand,omitempty"`
	TotalMobiles                 float64             `json:"totalMobiles"  bson:"totalMobiles,omitempty"`
	Village                      primitive.ObjectID  `json:"village"  bson:"village,omitempty"`
	YearlyIncome                 string              `json:"yearlyIncome"  bson:"yearlyIncome,omitempty"`
	FarmerOrg                    primitive.ObjectID  `json:"farmerOrg"  bson:"farmerOrg,omitempty"`
	FarmerID                     string              `json:"farmerID"  bson:"farmerID,omitempty"`
	DoorNumber                   string              `json:"doorNo"  bson:"doorNo,omitempty"`
	LandMark                     string              `json:"landMark"  bson:"landMark,omitempty"`
	Street                       string              `json:"street"  bson:"street,omitempty"`
	PreferredMarkets             []string            `json:"preferredMarkets"  bson:"preferredMarkets,omitempty"`
	IsDisabled                   bool                `json:"isDisabled"  bson:"isDisabled,omitempty"`
	FertilizerQuantity           float64             `json:"fertilizerQuantity"  bson:"fertilizerQuantity,omitempty"`
	RainDisasterFrom             float64             `json:"rainDisasterFrom"  bson:"rainDisasterFrom,omitempty"`
	RainDisasterTo               float64             `json:"rainDisasterTo"  bson:"rainDisasterTo,omitempty"`
	RainfallMedFrom              float64             `json:"rainfallMedFrom"  bson:"rainfallMedFrom,omitempty"`
	RainfallMedTo                float64             `json:"rainfallMedTo"  bson:"rainfallMedTo,omitempty"`
	RelativeHumidityDisasterFrom float64             `json:"relativeHumidityDisasterFrom"  bson:"relativeHumidityDisasterFrom,omitempty"`
	SeedQuantity                 float64             `json:"seedQuantity"  bson:"seedQuantity,omitempty"`
	TemperatureDisasterFrom      float64             `json:"temperatureDisasterFrom"  bson:"temperatureDisasterFrom,omitempty"`
	TemperatureDisasterTo        float64             `json:"temperatureDisasterTo"  bson:"temperatureDisasterTo,omitempty"`
	WindDirectionDisasterTo      float64             `json:"windDirectionDisasterTo"  bson:"windDirectionDisasterTo,omitempty"`
	WindDirectionMedFrom         float64             `json:"windDirectionMedFrom"  bson:"windDirectionMedFrom,omitempty"`
	RelativeHumidityDisasterTo   float64             `json:"relativeHumidityDisasterTo"  bson:"relativeHumidityDisasterTo,omitempty"`
	WindDirectionMedTo           float64             `json:"windDirectionMedTo"  bson:"windDirectionMedTo,omitempty"`
	WindSpeedDisasterFrom        float64             `json:"windSpeedDisasterFrom"  bson:"windSpeedDisasterFrom,omitempty"`
	WindSpeedDisasterTo          float64             `json:"windSpeedDisasterTo"  bson:"windSpeedDisasterTo,omitempty"`
	WindSpeedMedFrom             float64             `json:"windSpeedMedFrom"  bson:"windSpeedMedFrom,omitempty"`
	WindSpeedMedTo               float64             `json:"windSpeedMedTo"  bson:"windSpeedMedTo,omitempty"`
	UniqueId                     int                 `json:"uniqueId"  bson:"uniqueId,omitempty"`
	Password                     string              `json:"password"  bson:"password,omitempty"`
}

//RefFarmer : "Farmer with refrence data such as language..."
type RefFarmer struct {
	Farmer `bson:",inline"`
	Ref    struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//FarmerFilter : "Used for constructing filter query"
type FarmerFilter struct {
	State         []primitive.ObjectID `json:"state"  bson:"state,omitempty"`
	District      []primitive.ObjectID `json:"district"  bson:"district,omitempty"`
	Block         []primitive.ObjectID `json:"block"  bson:"block,omitempty"`
	GramPanchayat []primitive.ObjectID `json:"gramPanchayat"  bson:"gramPanchayat,omitempty"`
	Village       []primitive.ObjectID `json:"village"  bson:"village,omitempty"`
	SoilType      []primitive.ObjectID `json:"soilType"  bson:"soilType,omitempty"`
	FarmerOrg     []primitive.ObjectID `json:"farmerOrg"  bson:"farmerOrg,omitempty"`
	ActiveStatus  []bool               `json:"activeStatus" bson:"activeStatus,omitempty"`
	Status        []string             `json:"status" bson:"status,omitempty"`
	SortBy        string               `json:"sortBy"`
	SortOrder     int                  `json:"sortOrder"`
	Regex         struct {
		Name         string `json:"name" bson:"name"`
		MobileNumber string `json:"mobileNumber"  bson:"mobileNumber,omitempty"`
		SpouseName   string `json:"spouseName"  bson:"spouseName,omitempty"`
	} `json:"regex" bson:"regex"`
}
type CurrentLiveStocks struct {
	FarmerLiveStock primitive.ObjectID `json:"farmerLiveStock"  bson:"farmerLiveStock,omitempty"`
	Quantity        int                `json:"quantity"  bson:"quantity,omitempty"`
	Commodity       primitive.ObjectID `json:"commodity"  bson:"commodity,omitempty"`
	Category        primitive.ObjectID `json:"category"  bson:"category,omitempty"`
	Variety         primitive.ObjectID `json:"variety"  bson:"variety,omitempty"`
	Stage           primitive.ObjectID `json:"stage"  bson:"stage,omitempty"`
}
type CurrentCrops struct {
	FarmerCrops primitive.ObjectID `json:"farmerCrops"  bson:"farmerCrops,omitempty"`
	Commodity   primitive.ObjectID `json:"commodity"  bson:"commodity,omitempty"`
	Category    primitive.ObjectID `json:"category"  bson:"category,omitempty"`
	Variety     primitive.ObjectID `json:"variety"  bson:"variety,omitempty"`
	Season      primitive.ObjectID `json:"season"  bson:"season,omitempty"`
}
