package commentrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func EditComment(comment *models.Comment, db *gorm.DB) error {
	result := db.Save(&comment)

	return result.Error
}
