package postrepository

import (
	"SociLinkApi/models"
	authtypes "SociLinkApi/types/auth"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetPostsByUser(userId uuid.UUID, visibility authtypes.Visibility, db *gorm.DB) ([]models.Post, error) {
	var posts []models.Post

	result := db.Preload(clause.Associations).Where("user_id = ?", userId).Where("visibility IN ?", visibility.GetAllowedVisibilities()).Order("created_at desc").Find(&posts)

	return posts, result.Error
}
