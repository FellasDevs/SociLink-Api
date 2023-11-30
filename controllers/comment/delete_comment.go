package commentcontroller

import (
	commentrepository "SociLinkApi/repository/comment"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func DeleteComment(context *gin.Context, db *gorm.DB) {
	id := context.Param("id")
	commentId, err := uuid.Parse(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Id do comentário inválido",
		})
		return
	}

	if err = commentrepository.DeleteComment(commentId, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Erro ao deletar o comentário",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Comentário deletado com sucesso",
		})
	}
}
