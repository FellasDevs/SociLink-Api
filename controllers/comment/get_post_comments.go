package commentcontroller

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	commentrepository "SociLinkApi/repository/comment"
	postrepository "SociLinkApi/repository/post"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func GetPostComments(context *gin.Context, db *gorm.DB) {
	var params dto.GetPostCommentsRequestDto
	if err := context.ShouldBindQuery(&params); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Parâmetros inválidos",
		})
		return
	}

	postUuid, err := uuid.Parse(params.PostId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Id do post inválido",
		})
		return
	}

	var userId *uuid.UUID
	uid, exists := context.Get("userId")
	if exists {
		id := uid.(uuid.UUID)
		userId = &id
	}

	if err = postrepository.GetPost(&models.Post{ID: postUuid}, userId, db); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Post não encontrado",
		})
		return
	}

	if comments, err := commentrepository.GetPostComments(postUuid, params.PaginationRequestDto, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Erro ao buscar comentários do post",
		})
		return
	} else {
		response := dto.GetPostCommentsResponseDto{
			Comments: make([]dto.CommentResponseDto, len(comments)),
		}

		for i, comment := range comments {
			response.Comments[i] = dto.CommentToResponseDto(comment)
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    response,
		})
	}
}
