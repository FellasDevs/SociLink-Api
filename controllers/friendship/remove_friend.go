package friendshipcontroller

import (
	"SociLinkApi/models"
	friendshiprepository "SociLinkApi/repository/friendship"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func RemoveFriend(context *gin.Context, db *gorm.DB) {
	uid, _ := context.Get("userId")
	userId := uid.(uuid.UUID)

	friendshipId, err := uuid.Parse(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "id da amizade não informado",
		})
		return
	}

	friendship := models.Friendship{ID: friendshipId}
	if err = friendshiprepository.GetFriendship(&friendship, db); err != nil {
		var statusCode int

		if err.Error() == "record not found" {
			statusCode = http.StatusNotFound
		} else {
			statusCode = http.StatusInternalServerError
		}

		context.JSON(statusCode, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}

	if userId != friendship.UserID && (userId != friendship.FriendID || friendship.Pending) {
		context.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Você não tem permissão para remover esta amizade",
		})
		return
	}

	if err = friendshiprepository.DeleteFriendship(&friendship, db); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "amizade desfeita com sucesso",
		})
	}
}
