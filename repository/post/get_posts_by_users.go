package postrepository

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	authtypes "SociLinkApi/types/auth"
	"SociLinkApi/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetPostsByUsers(userIds []uuid.UUID, visibility authtypes.Visibility, pagination dto.PaginationRequestDto, db *gorm.DB) ([]models.Post, error) {
	var posts []models.Post

	query := db.Preload(clause.Associations)

	query = query.Where("user_id IN ?", userIds).Where("visibility IN ?", visibility.GetAllowedVisibilities())

	utils.UsePagination(query, pagination)

	result := query.Order("created_at desc").Find(&posts)

	return posts, result.Error
}
