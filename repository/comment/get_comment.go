package commentrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func GetComment(comment *models.Comment, db *gorm.DB) error {
	result := db.Preload("User").Find(&comment)

	return result.Error
}
