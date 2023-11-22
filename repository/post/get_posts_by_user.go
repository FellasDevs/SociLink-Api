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

func GetPostsByUser(userId uuid.UUID, visibility authtypes.Visibility, pagination dto.PaginationRequestDto, db *gorm.DB) ([]models.Post, error) {
	var posts []models.Post

	query := db.Preload(clause.Associations)

	query = query.Where("visibility IN ?", visibility.GetAllowedVisibilities()).Where("user_id = ?", userId)

	utils.UsePagination(query, pagination)

	result := query.Order("created_at desc").Find(&posts)

	return posts, result.Error
}
