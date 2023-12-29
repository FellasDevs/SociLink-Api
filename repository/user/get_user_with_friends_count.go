package userrepository

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	"gorm.io/gorm"
)

func GetUserWithFriendsCount(user models.User, db *gorm.DB) (dto.UserWithFriendsCount, error) {
	var userWithFriendsCount dto.UserWithFriendsCount

	query := db.Where(&user)

	query = query.Joins("LEFT JOIN friendships ON friendships.accepted = true AND (friendships.friend_id = users.id OR friendships.user_id = users.id)")

	query = query.Select("users.*, COUNT(friendships.id) AS friends_count").Group("users.id")

	result := query.First(&user).Scan(&userWithFriendsCount)

	return userWithFriendsCount, result.Error
}
