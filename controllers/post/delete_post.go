package postcontroller

import (
	"SociLinkApi/models"
	postrepository "SociLinkApi/repository/post"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeletePost(context *gin.Context, db *gorm.DB) {
	postIdString := context.Param("id")
	postId, err := uuid.Parse(postIdString)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID de post inválido",
		})
		return
	}

	uid, _ := context.Get("userId")
	userId := uid.(uuid.UUID)

	post := models.Post{ID: postId}
	if err = postrepository.GetPost(&post, &userId, db); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if post.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Você não possúi permissão para deletar este post",
		})
		return
	}

	post.Deleted = true

	if err = postrepository.UpdatePost(&post, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	context.JSON(http.StatusNoContent, gin.H{
		"success": true,
		"message": "Post deletado com sucesso",
	})
}
