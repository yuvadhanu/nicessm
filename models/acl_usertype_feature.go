package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//ACLUserTypeFeature : ""
type ACLUserTypeFeature struct {
	ID         primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID   string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	FeatureID  string             `json:"featureId" bson:"featureId,omitempty"`
	UserTypeID string             `json:"userTypeId" bson:"userTypeId,omitempty"`
	ModuleID   string             `json:"moduleId" bson:"moduleId,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Desc       string             `json:"desc" bson:"desc,omitempty"`
	Status     string             `json:"status" bson:"status,omitempty"`
	Created    Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated    []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
	Check      string             `json:"check"  bson:"check,omitempty"`
}

//RefACLUserTypeFeature : ""
type RefACLUserTypeFeature struct {
	ACLUserTypeFeature `bson:",inline"`
	Ref                struct {
		// Department *Department `json:"department,omitempty" bson:"department,omitempty"`
		// ULB        *ULB        `json:"ulb,omitempty" bson:"ulb,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//ACLUserTypeFeatureFilter : ""
type ACLUserTypeFeatureFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}

//UserTypeFeatureAccess : ""
type UserTypeFeatureAccess struct {
	UserType `bson:",inline"`
	Module   struct {
		Module   `bson:",inline"`
		Features []FeatureAccess `json:"features" bson:"features,omitempty"`
	} `json:"module" bson:"module,omitempty"`
}

//FeatureAccess : ""
type FeatureAccess struct {
	Feature `bson:",inline"`
	Access  *ACLUserTypeFeature `json:"access" bson:"access,omitempty"`
}
