package models

// Pagination : Pagination struct
type Pagination struct {
	PageNum   int `json:"pageNum" bson:"pageNum"`
	Limit     int `json:"limit" bson:"limit"`
	Count     int `json:"count" bson:"count"`
	NextPage  int `json:"nextPage" bson:"nextPage"`
	PrevPage  int `json:"prevPage" bson:"prevPage"`
	TotalPage int `json:"totalPage" bson:"totalPage"`
}
