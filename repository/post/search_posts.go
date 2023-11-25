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
		query = query.Joins("LEFT JOIN friendships ON (friendships.friend_id = posts.user_id OR friendships.user_id = posts.user_id)")

		query = query.Or("(visibility = ? OR visibility = ?) AND posts.user_id = ?", authtypes.Private, authtypes.Friends, userId)
		query = query.Or("visibility = ? AND EXISTS(SELECT * FROM friendships WHERE ((friendships.user_id = ? AND friendships.friend_id = posts.user_id) OR (friendships.friend_id = ? AND friendships.user_id = posts.user_id)) LIMIT 1)", authtypes.Friends, userId, userId)
	}

	utils.UsePagination(query, "*", &posts.PaginationResponse)

	result := query.Order("posts.created_at desc").Find(&posts.Posts).Scan(&posts.PaginationResponse)

	return posts, result.Error
}
