package userrepository

import (
	"SociLinkApi/dto"
	types "SociLinkApi/types/pagination"
	"SociLinkApi/utils"
	"gorm.io/gorm"
)

func GetUsersByNameOrNickname(search string, pagination dto.PaginationRequestDto, db *gorm.DB) (types.UserListing, error) {
	users := types.UserListing{
		PaginationResponse: types.PaginationResponse{
			Page:     pagination.Page,
			PageSize: pagination.PageSize,
		},
	}

	query := db.Where("name ILIKE ?", "%"+search+"%").Or("nickname ILIKE ?", "%"+search+"%")

	utils.UsePagination(query, "*", &users.PaginationResponse)

	result := query.Find(&users.Users).Scan(&users.PaginationResponse)

	return users, result.Error
}
