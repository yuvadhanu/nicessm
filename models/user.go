package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User : ""
type User struct {
	ID                           primitive.ObjectID   `json:"id" form:"id," bson:"_id,omitempty"`
	UserName                     string               `json:"userName" bson:"userName,omitempty"`
	Name                         string               `json:"name" bson:"name,omitempty"`
	FatherName                   string               `json:"fatherName" bson:"fatherName,omitempty"`
	SpouseName                   string               `json:"spouseName" bson:"spouseName,omitempty"`
	Gender                       string               `json:"gender" bson:"gender,omitempty"`
	Mobile                       string               `json:"mobile" bson:"mobile,omitempty"`
	Email                        string               `json:"email" bson:"email,omitempty"`
	DOB                          *time.Time           `json:"dob" bson:"dob,omitempty"`
	Address                      Address              `json:"address" bson:"address,omitempty"`
	Profile                      string               `json:"profile" bson:"profile,omitempty"`
	OrganisationID               string               `json:"organisationId" bson:"organisationId,omitempty"`
	Password                     string               `json:"-" bson:"password,omitempty"`
	Pass                         string               `json:"password" bson:"-"`
	Role                         string               `json:"role" bson:"role,omitempty"`
	Designation                  string               `json:"designation" bson:"designation,omitempty"`
	Type                         string               `json:"type" bson:"type,omitempty"`
	Created                      Created              `json:"createdOn" bson:"createdOn,omitempty"`
	Updated                      Updated              `json:"updated" form:"id," bson:"updated,omitempty"`
	UpdateLog                    []Updated            `json:"updatedLog" form:"id," bson:"updatedLog,omitempty"`
	Token                        string               `json:"token" bson:"-"`
	Status                       string               `json:"status" bson:"status,omitempty"`
	ManagerID                    string               `json:"managerId" bson:"managerId,omitempty"`
	BloodGroup                   string               `json:"bloodGroup" bson:"bloodGroup,omitempty"`
	CollectionLimit              CollectionLimit      `json:"collectionLimit" bson:"collectionLimit,omitempty"`
	AccessPrivilege              AccessPrivilege      `json:"accessPrivilege" bson:"accessPrivilege,omitempty"`
	Experience                   float64              `json:"experience" bson:"experience,omitempty"`
	EducationalQualification     string               `json:"educationalQualification" bson:"educationalQualification,omitempty"`
	KnowledgeDomains             []primitive.ObjectID `json:"knowledgeDomains" bson:"knowledgeDomains,omitempty"`
	SubDomains                   []primitive.ObjectID `json:"subDomains" bson:"subDomains,omitempty"`
	Organisation                 string               `json:"organisation" bson:"organisation,omitempty"`
	Officenumber                 int                  `json:"officenumber" bson:"officenumber,omitempty"`
	Officeaddress                string               `json:"officeddress " bson:"officeddress ,omitempty"`
	OrganisationNatureofBusiness string               `json:"organisationNatureofBusiness " bson:"organisationNatureofBusiness ,omitempty"`
	KdExpertise                  string               `json:"kdExpertise " bson:"kdExpertise ,omitempty"`
	LanguageExpertise            primitive.ObjectID   `json:"languageExpertise " bson:"languageExpertise"`
	SubjectExpertise             string               `json:"subjectExpertise " bson:"subjectExpertise ,omitempty"`
	Occupation                   string               `json:"occupation " bson:"occupation ,omitempty"`
}
type AccessPrivilege struct {
	AccessLevel string               `json:"accessLevel" bson:"accessLevel,omitempty"`
	Districts   []primitive.ObjectID `json:"districts" bson:"districts,omitempty"`
	States      []primitive.ObjectID `json:"states" bson:"states,omitempty"`
}

//RefUser :""
type RefUser struct {
	User `bson:",inline"`
	Ref  struct {
		Manager      User          `json:"manager" bson:"manager,omitempty"`
		Organisation *Organisation `json:"organisation" bson:"organisation,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//UserFilter : ""
type UserFilter struct {
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

//UserChangePassword : ""
type UserChangePassword struct {
	UserName    string `json:"userName" bson:"userName,omitempty"`
	OldPassword string `json:"oldPassword" bson:"oldPassword,omitempty"`
	NewPassword string `json:"newPassword" bson:"newPassword,omitempty"`
}

//UserLocation : ""
type UserLocation struct {
	UniqueID  string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	Status    string     `json:"status" bson:"status,omitempty"`
	UserName  string     `json:"userName" bson:"userName,omitempty"`
	Time      *time.Time `json:"time" bson:"time,omitempty"`
	Location  Location   `json:"location" bson:"location,omitempty"`
	UserType  string     `json:"userType" bson:"userType,omitempty"`
	Name      string     `json:"name" bson:"name,omitempty"`
	EntryType string     `json:"entryType" bson:"entryType,omitempty"`
	ErrMsg    string     `json:"errMsg" bson:"errMsg,omitempty"`
}

//RefUserLocation :""
type RefUserLocation struct {
	UserLocation `bson:",inline"`
	Ref          struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//UserLocationFilter : ""
type UserLocationFilter struct {
	Status   []string `json:"status"`
	UserType []string `json:"userType" bson:"userType,omitempty"`
}
type RefPassword struct {
	UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Token    string `json:"token" bson:"token,omitempty"`
	OTP      int64  `json:"otp" bson:"otp,omitempty"`
	Password string `json:"password" bson:"password,omitempty"`
}
type CollectionLimit struct {
	Cash float64 `json:"cash,omitempty" bson:"cash,omitempty"`
}
type UserStory struct {
	UniqueID string     `json:"uniqueId" bson:"uniqueId,omitempty"`
	UserName string     `json:"userName" bson:"userName,omitempty"`
	On       *time.Time `json:"on" bson:"on,omitempty"`
	By       string     `json:"by" bson:"by,omitempty"`
	ByType   string     `json:"byType" bson:"byType,omitempty"`
	Msg      string     `json:"msg" bson:"msg,omitempty"`
}
