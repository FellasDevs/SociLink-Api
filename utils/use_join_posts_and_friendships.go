package utils

import "gorm.io/gorm"

func UseJoinPostsAndFriendships(query *gorm.DB) {
	query = query.Joins("LEFT JOIN friendships ON (friendships.friend_id = posts.user_id OR friendships.user_id = posts.user_id)")
}
