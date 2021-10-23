package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//ACLUserTypeTab : ""
type ACLUserTypeTab struct {
	ID         primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	UniqueID   string             `json:"uniqueId" bson:"uniqueId,omitempty"`
	TabID      string             `json:"tabId" bson:"tabId,omitempty"`
	UserTypeID string             `json:"userTypeId" bson:"userTypeId,omitempty"`
	ModuleID   string             `json:"moduleId" bson:"moduleId,omitempty"`
	Name       string             `json:"name" bson:"name,omitempty"`
	Desc       string             `json:"desc" bson:"desc,omitempty"`
	Status     string             `json:"status" bson:"status,omitempty"`
	Created    Created            `json:"created,omitempty"  bson:"created,omitempty"`
	Updated    []Updated          `json:"updated,omitempty"  bson:"updated,omitempty"`
	Check      string             `json:"check"  bson:"check,omitempty"`
}

//RefACLUserTypeTab : ""
type RefACLUserTypeTab struct {
	ACLUserTypeTab `bson:",inline"`
	Ref            struct {
		// Department *Department `json:"department,omitempty" bson:"department,omitempty"`
		// ULB        *ULB        `json:"ulb,omitempty" bson:"ulb,omitempty"`
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}

//ACLUserTypeTabFilter : ""
type ACLUserTypeTabFilter struct {
	Status    []string `json:"status,omitempty" bson:"status,omitempty"`
	SortBy    string   `json:"sortBy,omitempty" bson:"sortBy,omitempty"`
	SortOrder int      `json:"sortOrder,omitempty" bson:"sortOrder,omitempty"`
}

//UserTypeTabAccess : ""
type UserTypeTabAccess struct {
	UserType `bson:",inline"`
	Module   struct {
		Module `bson:",inline"`
		Tabs   []TabAccess `json:"tabs" bson:"tabs,omitempty"`
	} `json:"module" bson:"module,omitempty"`
}

//TabAccess : ""
type TabAccess struct {
	Tab    `bson:",inline"`
	Access *ACLUserTypeTab `json:"access" bson:"access,omitempty"`
}
