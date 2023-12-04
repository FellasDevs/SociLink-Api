package commentrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateComment(comment *models.Comment, db *gorm.DB) error {
	result := db.Preload("User").Clauses(clause.Returning{}).Create(&comment)

	return result.Error
}
