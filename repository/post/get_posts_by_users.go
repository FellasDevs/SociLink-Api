package postrepository

import (
	"SociLinkApi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetPostsByUsers(userIds []uuid.UUID, db *gorm.DB) ([]models.Post, error) {
	var posts []models.Post

	result := db.Preload(clause.Associations).Where("user_id IN ?", userIds).Order("created_at").Find(&posts)

	return posts, result.Error
}
