package userroutes

import (
	usercontroller "SociLinkApi/controllers/user"
	"SociLinkApi/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

type GetUserParams struct {
	id uuid.UUID `uri:"string" binding:"required, string"`
}

func GetUserRoute(router *gin.Engine, db *gorm.DB) {
	router.GET("/users/:id", func(context *gin.Context) {
		var params GetUserParams

		if err := context.ShouldBindUri(&params); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err,
			})
			return
		}

		user, err := usercontroller.GetUser(params.id, db)

		if err != nil {
			utils.SendUnknownError(err, context)
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "",
			"user":    user,
		})
	})
}
