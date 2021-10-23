package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//ACLUserTypeModule : ""
type ACLUserTypeModule struct {
	ID         primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID   string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	ModuleID   string             `json:"moduleId" bson:"moduleId,omitempty"`
	UserTypeID string             `json:"userTypeId" bson:"userTypeId,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Desc       string             `json:"desc" bson:"desc,omitempty"`
	Status     string             `json:"status" bson:"status,omitempty"`
	Created    Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated    []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
	Check      string             `json:"check"  bson:"check,omitempty"`
}

//RefACLUserTypeModule : ""
type RefACLUserTypeModule struct {
	ACLUserTypeModule `bson:",inline"`
	Ref               struct {
		// Department *Department `json:"department,omitempty" bson:"department,omitempty"`
		// ULB        *ULB        `json:"ulb,omitempty" bson:"ulb,omitempty"`
		Module *Module `json:"module,omitempty" bson:"module,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//ACLUserTypeModuleFilter : ""
type ACLUserTypeModuleFilter struct {
	Module    []string `json:"module,omitempty" bson:"module,omitempty"`
	UserType  []string `json:"userType,omitempty" bson:"userType,omitempty"`
	Check     []string `json:"check,omitempty" bson:"check,omitempty"`
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}

//UserTypeModuleAccess : ""
type UserTypeModuleAccess struct {
	UserType `bson:",inline"`
	Modules  []ModuleAccess `json:"modules" bson:"modules,omitempty"`
}

//ModuleAccess : ""
type ModuleAccess struct {
	Module `bson:",inline"`
	Access *ACLUserTypeModule `json:"access" bson:"access,omitempty"`
}
