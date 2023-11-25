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
	postIDString := context.Param("id")
	postID, err := uuid.Parse(postIDString)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID de post inválido",
		})
		return
	}

	uid, _ := context.Get("userId")
	userId := uid.(uuid.UUID)

	post := models.Post{ID: postID}
	if err = postrepository.GetPost(&post, &userId, db); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
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
