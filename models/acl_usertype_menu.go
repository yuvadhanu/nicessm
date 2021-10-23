package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//ACLUserTypeMenu : ""
type ACLUserTypeMenu struct {
	ID         primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID   string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	MenuID     string             `json:"menuId" bson:"menuId,omitempty"`
	UserTypeID string             `json:"userTypeId" bson:"userTypeId,omitempty"`
	ModuleID   string             `json:"moduleId" bson:"moduleId,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Desc       string             `json:"desc" bson:"desc,omitempty"`
	Status     string             `json:"status" bson:"status,omitempty"`
	Created    Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated    []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
	Check      string             `json:"check"  bson:"check,omitempty"`
}

//RefACLUserTypeMenu : ""
type RefACLUserTypeMenu struct {
	ACLUserTypeMenu `bson:",inline"`
	Ref             struct {
		// Department *Department `json:"department,omitempty" bson:"department,omitempty"`
		// ULB        *ULB        `json:"ulb,omitempty" bson:"ulb,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//ACLUserTypeMenuFilter : ""
type ACLUserTypeMenuFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}

//UserTypeMenuAccess : ""
type UserTypeMenuAccess struct {
	UserType `bson:",inline"`
	Module   struct {
		Module `bson:",inline"`
		Menus  []MenuAccess `json:"menus" bson:"menus,omitempty"`
	} `json:"module" bson:"module,omitempty"`
}

//MenuAccess : ""
type MenuAccess struct {
	Menu   `bson:",inline"`
	Access *ACLUserTypeMenu `json:"access" bson:"access,omitempty"`
}
