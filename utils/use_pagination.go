package utils

import (
	types "SociLinkApi/types/pagination"
	"gorm.io/gorm"
)

func UsePagination(query *gorm.DB, selectArgs string, pagination *types.PaginationResponse) {
	if pagination.Page == 0 {
		pagination.Page = 1
	}

	if pagination.PageSize == 0 {
		pagination.PageSize = 50
	}

	query = query.Select(selectArgs + ", COUNT(*) OVER() AS total_count")

	query = query.Limit(pagination.PageSize).Offset(pagination.PageSize * (pagination.Page - 1))
}
