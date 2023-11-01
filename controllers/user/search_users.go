package usercontroller

import (
	"SociLinkApi/models"
	userrepository "SociLinkApi/repository/user"
	"gorm.io/gorm"
)

func SearchUsers(search string, db *gorm.DB) ([]models.User, error) {
	return userrepository.GetUsersByName(search, db)
}
