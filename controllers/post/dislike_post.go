package postcontroller

import (
	"SociLinkApi/models"
	likerepository "SociLinkApi/repository/like"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func DislikePost(context *gin.Context, db *gorm.DB) {
	uid, _ := context.Get("userId")
	userId := uid.(uuid.UUID)

	postId := context.Param("id")
	if postId == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Id do post não informado",
		})
		return
	}

	postUuid, err := uuid.Parse(postId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Id do post inválido",
		})
		return
	}

	like := models.Like{
		UserID: userId,
		PostID: postUuid,
	}

	if err = likerepository.DeleteLike(&like, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Post descurtido com sucesso",
		})
	}
}
