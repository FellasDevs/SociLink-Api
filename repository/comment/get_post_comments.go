package commentrepository

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	"SociLinkApi/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetPostComments(postId uuid.UUID, pagination dto.PaginationRequestDto, db *gorm.DB) ([]models.Comment, error) {
	var comments []models.Comment

	query := db.Preload("User").Where("post_id = ?", postId)

	utils.UsePagination(query, pagination)

	result := query.Order("created_at desc").Find(&comments)

	return comments, result.Error
}
