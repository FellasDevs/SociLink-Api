package commentrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func CreateComment(comment *models.Comment, db *gorm.DB) error {
	result := db.Create(&comment)

	return result.Error
}
