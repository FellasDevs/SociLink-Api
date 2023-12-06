package postcontroller

import (
	"SociLinkApi/models"
	postrepository "SociLinkApi/repository/post"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func RecoverPost(context *gin.Context, db *gorm.DB) {
	postId := context.Param("id")

	postUuid, err := uuid.Parse(postId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Id do post inválido",
		})
		return
	}

	uid, _ := context.Get("userId")
	userId := uid.(uuid.UUID)

	post := models.Post{
		ID:      postUuid,
		Deleted: true,
	}

	if err = postrepository.GetDeletedPost(&post, &userId, db); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Post não encontrado",
		})
		return
	}

	if post.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Você não tem permissão para recuperar este post",
		})
		return
	}

	post.Deleted = false

	if err = postrepository.UpdatePost(&post, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Erro ao recuperar post",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Post recuperado com sucesso",
	})
}
