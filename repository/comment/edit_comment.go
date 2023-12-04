package commentrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func EditComment(comment *models.Comment, db *gorm.DB) error {
	result := db.Preload("User").Clauses(clause.Returning{}).Save(&comment)

	return result.Error
}
