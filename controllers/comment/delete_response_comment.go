package commentcontroller

import (
	commentrepository "SociLinkApi/repository/comment"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeleteCommentReply(context *gin.Context, db *gorm.DB) {
	replyId := context.Param("id")
	replyUuid, err := uuid.Parse(replyId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Id da resposta do comentário inválido",
		})
		return
	}

	if err = commentrepository.DeleteCommentReply(replyUuid, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Erro ao excluir a resposta do comentário",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Resposta do comentário excluída com sucesso",
		})
	}
}
