package userrepository

import (
	"SociLinkApi/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserById(id uuid.UUID, db *gorm.DB) (models.User, error) {
	var user models.User

	result := db.Preload("Users").First(&user, id)

	return user, result.Error
}
