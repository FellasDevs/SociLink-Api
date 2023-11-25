package timelinerepository

import (
	authtypes "SociLinkApi/types/auth"
	types "SociLinkApi/types/pagination"
	"SociLinkApi/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserTimeline(userId *uuid.UUID, timelineUserId uuid.UUID, posts *types.PostListing, db *gorm.DB) error {
	query := db.Preload("User")

	utils.UseJoinPostsAndFriendships(query)

	query = query.Where("posts.user_id = ?", timelineUserId)
	query = query.Where("visibility = ?", authtypes.Public)

	query = query.Or("posts.user_id = ?", timelineUserId)
	query = query.Where("visibility = ?", authtypes.Friends)
	utils.UseAreUsersFriends(query, *userId, timelineUserId)

	utils.UsePagination(query, "posts.id, posts.content, posts.images, posts.visibility, posts.user_id, posts.created_at", &posts.PaginationResponse)

	result := query.Order("created_at desc").Find(&posts.Posts).Scan(&posts.PaginationResponse)

	return result.Error
}
