package commentcontroller

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	commentrepository "SociLinkApi/repository/comment"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func EditComment(context *gin.Context, db *gorm.DB) {
	commentId := context.Param("id")
	if commentId == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Id do comentário não informado",
		})
		return
	}

	commentUuid, err := uuid.Parse(commentId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Id do comentário inválido",
		})
		return
	}

	var params dto.EditCommentRequestDto
	if err = context.ShouldBindJSON(&params); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Parâmetros inválidos",
		})
		return
	}

	comment := models.Comment{ID: commentUuid}
	if err = commentrepository.GetComment(&comment, db); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Comentário não encontrado",
		})
		return
	}

	comment.Content = params.Content

	if err = commentrepository.EditComment(&comment, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Erro ao editar o comentário",
		})
	} else {
		response := dto.EditCommentResponseDto{
			Comment: dto.CommentToResponseDto(comment),
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Comentário editado com sucesso",
			"data":    response,
		})
	}
}
