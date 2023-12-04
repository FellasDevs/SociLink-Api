package commentcontroller

import (
	"SociLinkApi/models"
	likerepository "SociLinkApi/repository/like"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateCommentLike(context *gin.Context, db *gorm.DB) {
	uid, _ := context.Get("userId")
	userId := uid.(uuid.UUID)

	commentId := context.Param("commentId")
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

	like := models.CommentLike{
		UserID:    userId,
		CommentID: commentUuid,
	}

	err = likerepository.CreateCommentLike(&like, db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Erro ao curtir o comentário",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Curtida no comentário adicionada com sucesso",
	})
}
