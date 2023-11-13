package postrepository

import (
	"SociLinkApi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetPostsByUser(userId uuid.UUID, db *gorm.DB) ([]models.Post, error) {
	var posts []models.Post

	result := db.Preload(clause.Associations).Where("user_id = ?", userId).Order("created_at").Find(&posts)

	return posts, result.Error
}
