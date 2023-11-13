package userrepository

import (
	"SociLinkApi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetUserById(id uuid.UUID, db *gorm.DB) (models.User, error) {
	var user models.User

	result := db.Preload(clause.Associations).First(&user, id)

	return user, result.Error
}
