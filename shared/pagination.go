package shared

import "nicessm-api-service/models"

//PaginationData : ""
func (s *Shared) PaginationData(pagination *models.Pagination) {
	if pagination.Limit > 0 {
		pagination.TotalPage = (pagination.Count + pagination.Limit - 1) / pagination.Limit
	}
	if (pagination.PageNum) < pagination.TotalPage {
		pagination.NextPage = pagination.PageNum + 1
	} else {
		pagination.NextPage = -1
	}
	if (pagination.PageNum) > 1 {
		pagination.PrevPage = pagination.PageNum - 1
	} else {
		pagination.PrevPage = -1
	}
	return
}
