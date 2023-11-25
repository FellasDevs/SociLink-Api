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
	uid, exist := context.Get("userId")
	var userId *uuid.UUID
	if exist {
		id := uid.(uuid.UUID)
		userId = &id
	}

	var params dto.SearchPostRequestDto
	if err := context.ShouldBindQuery(&params); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if posts, err := postrepository.SearchPosts(params.Search, userId, params.PaginationRequestDto, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		response := dto.SearchPostResponseDto{
			PaginationResponse: posts.PaginationResponse,
			Posts:              make([]dto.PostResponseDto, len(posts.Posts)),
		}

		for i, post := range posts.Posts {
			response.Posts[i] = dto.PostToPostResponseDto(post)
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "posts encontrados",
			"data":    response,
		})
	}
}
