package postrepository

import (
	"SociLinkApi/models"
	authtypes "SociLinkApi/types/auth"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SearchPosts(search string, userId *uuid.UUID, db *gorm.DB) ([]models.Post, error) {
	var posts []models.Post

	query := db.Preload("User")

	query = query.Where("content ILIKE ?", "%"+search+"%")

	query = query.Where("visibility = ?", authtypes.Public)

	if userId != nil {
		query = query.Joins("LEFT JOIN friendships ON (friendships.friend_id = posts.user_id OR friendships.user_id = posts.user_id)")

		query = query.Or("(visibility = ? OR visibility = ?) AND posts.user_id = ?", authtypes.Private, authtypes.Friends, userId)
		query = query.Or("visibility = ? AND EXISTS(SELECT * FROM friendships WHERE ((friendships.user_id = ? AND friendships.friend_id = posts.user_id) OR (friendships.friend_id = ? AND friendships.user_id = posts.user_id)) LIMIT 1)", authtypes.Friends, userId, userId)
	}

	result := query.Order("created_at desc").Find(&posts)

	return posts, result.Error
}
