package postrepository

import (
	"SociLinkApi/dto"
	authtypes "SociLinkApi/types/auth"
	types "SociLinkApi/types/pagination"
	"SociLinkApi/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SearchPosts(search string, userId *uuid.UUID, pagination dto.PaginationRequestDto, db *gorm.DB) (types.PostListing, error) {
	posts := types.PostListing{
		PaginationResponse: types.PaginationResponse{
			Page:     pagination.Page,
			PageSize: pagination.PageSize,
		},
	}

	query := db.Preload("User")

	query = query.Where("content ILIKE ?", "%"+search+"%")
	query = query.Where("visibility = ?", authtypes.Public)

	if userId != nil {
		utils.UseJoinPostsAndFriendships(query)

		query = query.Or("content ILIKE ?", "%"+search+"%")
		query = query.Where("(visibility = ? OR visibility = ?) AND posts.user_id = ?", authtypes.Private, authtypes.Friends, userId)

		query = query.Or("content ILIKE ?", "%"+search+"%")
		query = query.Where("visibility = ?", authtypes.Friends)
		utils.UseAreUserAndPostOwnerFriends(query, *userId)
	}

	utils.UsePagination(query, "DISTINCT posts.id, posts.content, posts.images, posts.visibility, posts.user_id, posts.created_at", &posts.PaginationResponse)

	result := query.Order("posts.created_at desc").Find(&posts.Posts).Scan(&posts.PaginationResponse)

	return posts, result.Error
}
