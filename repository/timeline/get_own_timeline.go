package timelinerepository

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	authtypes "SociLinkApi/types/auth"
	"SociLinkApi/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetOwnTimeline(userId uuid.UUID, pagination dto.PaginationRequestDto, db *gorm.DB) ([]models.Post, error) {
	var posts []models.Post

	query := db.Preload("User")

	utils.UseJoinPostsAndFriendships(query)

	query = query.Where("posts.user_id = ?", userId)

	query = query.Or("visibility = ? OR visibility = ?", authtypes.Friends, authtypes.Public)
	utils.UseAreUserAndPostOwnerFriends(query, userId)

	utils.UsePagination(query, pagination)

	query = query.Select("DISTINCT posts.id, posts.content, posts.images, posts.visibility, posts.user_id, posts.created_at")

	result := query.Order("posts.created_at desc").Find(&posts)

	return posts, result.Error
}
