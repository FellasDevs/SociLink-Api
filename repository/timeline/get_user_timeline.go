package timelinerepository

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	authtypes "SociLinkApi/types/auth"
	"SociLinkApi/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserTimeline(userId *uuid.UUID, timelineUserId uuid.UUID, pagination dto.PaginationRequestDto, db *gorm.DB) ([]models.Post, error) {
	var posts []models.Post

	query := db.Preload("User")

	utils.UseJoinPostsAndFriendships(query)

	query = query.Where("posts.user_id = ?", timelineUserId)
	query = query.Where("visibility = ?", authtypes.Public)

	query = query.Or("posts.user_id = ?", timelineUserId)
	query = query.Where("visibility = ?", authtypes.Friends)
	utils.UseAreUsersFriends(query, *userId, timelineUserId)

	utils.UsePagination(query, pagination)

	query = query.Select("posts.id, posts.content, posts.images, posts.visibility, posts.user_id, posts.created_at")

	result := query.Order("created_at desc").Find(&posts)

	return posts, result.Error
}
