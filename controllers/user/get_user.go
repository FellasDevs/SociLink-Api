package usercontroller

import (
	"SociLinkApi/models"
	userrepository "SociLinkApi/repository/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUser(id uuid.UUID, db *gorm.DB) (models.User, error) {
	return userrepository.GetUserById(id, db)
}
