package postrepository

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	"SociLinkApi/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetPostsByUserId(userId uuid.UUID, pagination dto.PaginationRequestDto, db *gorm.DB) ([]models.Post, error) {
	var posts []models.Post

	query := db.Preload("User")

	query = query.Where("user_id = ?", userId).Where("deleted = ?", false)

	utils.UsePagination(query, pagination)

	query = query.Select("posts.id, posts.content, posts.images, posts.visibility, posts.user_id, posts.created_at")

	result := query.Order("created_at desc").Find(&posts)

	return posts, result.Error
}
