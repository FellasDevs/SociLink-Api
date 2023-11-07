package postcontroller

import (
	postrepository "SociLinkApi/repository/post"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeletePost(context *gin.Context, db *gorm.DB) {
	postIDString := context.Param("id")

	postID, err := uuid.Parse(postIDString)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID de post inválido",
		})
		return
	}

	post, err := postrepository.GetPost(postID, db)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	uid, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Token de autentificação não encontrada",
		})
		return
	}

	userId := uid.(uuid.UUID)

	if userId != post.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Você não tem permissão para deletar este post",
		})
		return
	}

	err = postrepository.DeletePost(postID, db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Post deletado com sucesso",
	})
}
