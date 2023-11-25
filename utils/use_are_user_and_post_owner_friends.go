package utils

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UseAreUserAndPostOwnerFriends(query *gorm.DB, userId uuid.UUID) {
	query = query.Where("EXISTS(SELECT * FROM friendships WHERE friendships.accepted = true AND ((friendships.user_id = ? AND friendships.friend_id = posts.user_id) OR (friendships.friend_id = ? AND friendships.user_id = posts.user_id)) LIMIT 1)", userId, userId)
}
