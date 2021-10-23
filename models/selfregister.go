package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User : ""
type SelfRegister struct {
	ID              primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	UserName        string             `json:"userName" bson:"userName,omitempty"`
	Name            string             `json:"name" bson:"name,omitempty"`
	FatherName      string             `json:"fatherName" bson:"fatherName,omitempty"`
	SpouseName      string             `json:"spouseName" bson:"spouseName,omitempty"`
	Gender          string             `json:"gender" bson:"gender,omitempty"`
	Mobile          string             `json:"mobile" bson:"mobile,omitempty"`
	Email           string             `json:"email" bson:"email,omitempty"`
	DOB             *time.Time         `json:"dob" bson:"dob,omitempty"`
	Address         Address            `json:"address" bson:"address,omitempty"`
	Profile         string             `json:"profile" bson:"profile,omitempty"`
	OrganisationID  string             `json:"organisationId" bson:"organisationId,omitempty"`
	Password        string             `json:"-" bson:"password,omitempty"`
	Pass            string             `json:"password" bson:"-"`
	Role            string             `json:"role" bson:"role,omitempty"`
	Designation     string             `json:"designation" bson:"designation,omitempty"`
	Type            string             `json:"type" bson:"type,omitempty"`
	Created         Created            `json:"createdOn" bson:"createdOn,omitempty"`
	Updated         Updated            `json:"updated" form:"id," bson:"updated,omitempty"`
	UpdateLog       []Updated          `json:"updatedLog" form:"id," bson:"updatedLog,omitempty"`
	Token           string             `json:"token" bson:"-"`
	Status          string             `json:"status" bson:"status,omitempty"`
	ManagerID       string             `json:"managerId" bson:"managerId,omitempty"`
	BloodGroup      string             `json:"bloodGroup" bson:"bloodGroup,omitempty"`
	CollectionLimit CollectionLimit    `json:"collectionLimit" bson:"collectionLimit,omitempty"`
}

//RefUser :""
type RefSelfRegister struct {
	SelfRegister `bson:",inline"`
	Ref          struct {
		Manager      SelfRegister  `json:"manager" bson:"manager,omitempty"`
		Organisation *Organisation `json:"organisation" bson:"organisation,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//UserFilter : ""
type SelfRegisterFilter struct {
	Status            []string `json:"status"`
	UniqueID          []string `json:"uniqueId"`
	OmitID            []string `json:"omitId"`
	OrganisationID    []string `json:"organisationId" bson:"organisationId,omitempty"`
	Manager           []string `json:"manager" bson:"manager,omitempty"`
	Type              []string `json:"type" bson:"type"`
	GetRecentLocation bool     `json:"getRecentLocation"`
	Regex             struct {
		Name     string `json:"name" bson:"name"`
		Contact  string `json:"contact" bson:"contact"`
		UserName string `json:"userName" bson:"userName"`
	} `json:"regex" bson:"regex"`
	SortBy    string `json:"sortBy"`
	SortOrder int    `json:"sortOrder"`
}
