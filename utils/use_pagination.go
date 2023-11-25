package utils

import (
	"SociLinkApi/dto"
	"gorm.io/gorm"
)

func UsePagination(query *gorm.DB, pagination dto.PaginationRequestDto) {
	if pagination.Page == 0 {
		pagination.Page = 1
	}

	if pagination.PageSize == 0 {
		pagination.PageSize = 30
	}

	query = query.Limit(pagination.PageSize).Offset(pagination.PageSize * (pagination.Page - 1))
}
