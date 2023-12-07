package postrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetPosts(postModel models.Post, db *gorm.DB) ([]models.Post, error) {
	var posts []models.Post

	result := db.Preload(clause.Associations).Where(postModel).Order("created_at desc").Find(&posts)

	return posts, result.Error
}
