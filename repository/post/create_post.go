package postrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm/clause"

	"gorm.io/gorm"
)

func CreatePost(post *models.Post, db *gorm.DB) error {
	result := db.Preload(clause.Associations).Clauses(clause.Returning{}).Create(&post)

	var user models.User
	db.Where(models.User{ID: post.UserID}).First(&user)
	post.User = user

	return result.Error
}
