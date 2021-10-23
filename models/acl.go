package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//UserModuleAccess : ""
type UserModuleAccess struct {
	ID         primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID   string             `json:"userName" bson:"userName,omitempty"`
	ModuleID   string             `json:"moduleId" bson:"moduleId,omitempty"`
	UserTypeID string             `json:"userTypeId" bson:"userTypeId,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Desc       string             `json:"desc" bson:"desc,omitempty"`
	Status     string             `json:"status" bson:"status,omitempty"`
	Created    Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated    []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
}

//UserMenuAccess : ""
type UserMenuAccess struct {
	ID         primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID   string             `json:"userName" bson:"userName,omitempty"`
	ModuleID   string             `json:"moduleId" bson:"moduleId,omitempty"`
	UserTypeID string             `json:"userTypeId" bson:"userTypeId,omitempty"`
	MenuID     string             `json:"menuId" bson:"menuId,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Desc       string             `json:"desc" bson:"desc,omitempty"`
	Status     string             `json:"status" bson:"status,omitempty"`
	Created    Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated    []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
}

//UserTabAccess : ""
type UserTabAccess struct {
	ID         primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID   string             `json:"userName" bson:"userName,omitempty"`
	ModuleID   string             `json:"moduleId" bson:"moduleId,omitempty"`
	UserTypeID string             `json:"userTypeId" bson:"userTypeId,omitempty"`
	TabID      string             `json:"tabId" bson:"tabId,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Desc       string             `json:"desc" bson:"desc,omitempty"`
	Status     string             `json:"status" bson:"status,omitempty"`
	Created    Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated    []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
}

//UserFeatureAccess : ""
type UserFeatureAccess struct {
	ID         primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID   string             `json:"userName" bson:"userName,omitempty"`
	ModuleID   string             `json:"moduleId" bson:"moduleId,omitempty"`
	UserTypeID string             `json:"userTypeId" bson:"userTypeId,omitempty"`
	FeatureID  string             `json:"featureId" bson:"featureId,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Desc       string             `json:"desc" bson:"desc,omitempty"`
	Status     string             `json:"status" bson:"status,omitempty"`
	Created    Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated    []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
}

//ACLAccess : ""
type ACLAccess struct {
	ModuleAccess []ModuleAccess
	// Module       *UserTypeModuleAccess
	MenuAccess []MenuAccess
	// Menu          []UserTypeMenuAccess
	// Tab           []UserTypeTabAccess
	TabAccess []TabAccess
	// Feature       []UserTypeFeatureAccess
	FeatureAccess []FeatureAccess
}
