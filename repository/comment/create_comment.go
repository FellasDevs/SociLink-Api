package commentrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateComment(comment *models.Comment, db *gorm.DB) error {
	result := db.Preload("User").Clauses(clause.Returning{}).Create(&comment)

	var user models.User
	db.Where(&models.User{ID: comment.UserID}).First(&user)
	comment.User = user

	return result.Error
}
