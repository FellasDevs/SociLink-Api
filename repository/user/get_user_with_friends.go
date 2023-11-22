package userrepository

import (
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func GetUserWithFriends(user *models.User, db *gorm.DB) error {
	query := db.Preload("Friends").Where(&user)

	query = query.Joins("LEFT JOIN friendships ON (friendships.friend_id = users.id OR friendships.user_id = users.id)")

	result := query.First(&user)

	return result.Error
}
