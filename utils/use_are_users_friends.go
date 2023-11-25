package utils

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UseAreUsersFriends(query *gorm.DB, user1 uuid.UUID, user2 uuid.UUID) {
	query = query.Where("EXISTS(SELECT * FROM friendships WHERE friendships.accepted = true AND ((friendships.user_id = ? AND friendships.friend_id = ?) OR (friendships.friend_id = ? AND friendships.user_id = ?)) LIMIT 1)", user1, user2, user1, user2)
}
