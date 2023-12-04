package commentcontroller

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	commentrepository "SociLinkApi/repository/comment"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateCommentReply(context *gin.Context, db *gorm.DB) {
	uid, _ := context.Get("userId")
	userId := uid.(uuid.UUID)

	var params dto.CreateCommentReplyRequestDto
	if err := context.ShouldBindJSON(&params); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Parâmetros inválidos",
		})
		return
	}

	commentId, err := uuid.Parse(params.CommentId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Id do comentário inválido",
		})
		return
	}

	var parentComment models.Comment
	if err := commentrepository.GetComment(&parentComment, db); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Comentário pai não encontrado",
		})
		return
	}

	reply := models.CommentReply{
		UserID:    userId,
		CommentID: commentId,
		Content:   params.Content,
		CreatedAt: time.Now(),
	}

	if err := commentrepository.CreateCommentReply(&reply, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Erro ao adicionar a resposta do comentário",
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Resposta do comentário adicionada com sucesso",
	})
}
