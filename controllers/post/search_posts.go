package postcontroller

import (
	postrepository "SociLinkApi/repository/post"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func SearchPosts(context *gin.Context, db *gorm.DB) {
	search := context.Param("search")
	if search == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "pesquisa inválida",
		})
		return
	}

	uid, exist := context.Get("userId")
	var userId *uuid.UUID
	if exist {
		id := uid.(uuid.UUID)
		userId = &id
	}

	if posts, err := postrepository.SearchPosts(search, userId, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "posts encontrados",
			"data":    posts,
		})
	}
}
