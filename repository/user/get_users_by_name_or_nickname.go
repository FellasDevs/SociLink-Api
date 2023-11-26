package userrepository

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	"SociLinkApi/utils"
	"gorm.io/gorm"
)

func GetUsersByNameOrNickname(search string, pagination dto.PaginationRequestDto, db *gorm.DB) ([]models.User, error) {
	var users []models.User

	query := db.Where("name ILIKE ?", "%"+search+"%").Or("nickname ILIKE ?", "%"+search+"%")

	utils.UsePagination(query, pagination)

	result := query.Find(&users)

	return users, result.Error
}
