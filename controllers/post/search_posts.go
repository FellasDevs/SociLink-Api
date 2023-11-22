package postcontroller

import (
	"SociLinkApi/dto"
	postrepository "SociLinkApi/repository/post"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func SearchPosts(context *gin.Context, db *gorm.DB) {
	search := context.Param("search")
	if search == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "pesquisa inválida",
		})
		return
	}

	uid, exist := context.Get("userId")
	var userId *uuid.UUID
	if exist {
		id := uid.(uuid.UUID)
		userId = &id
	}

	var pagination dto.PaginationRequestDto
	if err := context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if posts, err := postrepository.SearchPosts(search, userId, pagination, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		response := dto.SearchPostResponseDto{
			Posts: make([]dto.PostResponseDto, len(posts)),
		}

		for i, post := range posts {
			response.Posts[i] = dto.PostToPostResponseDto(post)
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "posts encontrados",
			"data":    response,
		})
	}
}
