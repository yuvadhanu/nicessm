package models

//Month : ""
type Month struct {
	Name    string `json:"name,omitempty" bson:"name,omitempty"`
	FYOrder int    `json:"fyOrder" bson:"fyOrder,omitempty"`
	Month   int    `json:"month" bson:"month,omitempty"`
}
