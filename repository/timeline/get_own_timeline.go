package timelinerepository

import (
	"SociLinkApi/dto"
	authtypes "SociLinkApi/types/auth"
	types "SociLinkApi/types/pagination"
	"SociLinkApi/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetOwnTimeline(userId uuid.UUID, pagination dto.PaginationRequestDto, db *gorm.DB) (types.PostListing, error) {
	posts := types.PostListing{
		PaginationResponse: types.PaginationResponse{
			Page:     pagination.Page,
			PageSize: pagination.PageSize,
		},
	}

	query := db.Preload("User")

	utils.UseJoinPostsAndFriendships(query)

	query = query.Where("posts.user_id = ?", userId)

	query = query.Or("visibility = ? OR visibility = ?", authtypes.Friends, authtypes.Public)
	utils.UseAreUserAndPostOwnerFriends(query, userId)

	utils.UsePagination(query, "DISTINCT posts.id, posts.content, posts.images, posts.visibility, posts.user_id, posts.created_at", &posts.PaginationResponse)

	result := query.Order("posts.created_at desc").Find(&posts.Posts).Scan(&posts.PaginationResponse)

	return posts, result.Error
}
