package postrepository

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	authtypes "SociLinkApi/types/auth"
	"SociLinkApi/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SearchPosts(search string, userId *uuid.UUID, pagination dto.PaginationRequestDto, db *gorm.DB) ([]models.Post, error) {
	var posts []models.Post

	query := db.Preload("User")

	query = query.Where("content ILIKE ?", "%"+search+"%")
	query = query.Where("visibility = ?", authtypes.Public)
	query = query.Where("deleted = ?", false)

	if userId != nil {
		utils.UseJoinPostsAndFriendships(query)

		query = query.Or("content ILIKE ?", "%"+search+"%")
		query = query.Where("(visibility = ? OR visibility = ?) AND posts.user_id = ?", authtypes.Private, authtypes.Friends, userId)
		query = query.Where("deleted = ?", false)

		query = query.Or("content ILIKE ?", "%"+search+"%")
		query = query.Where("visibility = ?", authtypes.Friends)
		query = query.Where("deleted = ?", false)
		utils.UseAreUserAndPostOwnerFriends(query, *userId)
	}

	utils.UsePagination(query, pagination)

	query = query.Select("posts.id, posts.original_post_id, posts.content, posts.images, posts.visibility, posts.user_id, posts.created_at, posts.deleted")

	result := query.Order("posts.created_at desc").Find(&posts)

	return posts, result.Error
}
