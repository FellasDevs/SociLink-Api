package middlewares

import (
	userrepository "SociLinkApi/repository/user"
	authservice "SociLinkApi/services/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type authHeader struct {
	AuthToken string `header:"Authorization"`
}

func AuthenticateUser(context *gin.Context, db *gorm.DB) {
	header := authHeader{}

	if err := context.ShouldBindHeader(&header); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Token de autorização não informado.",
		})

		context.Abort()
		return
	}

	token := strings.Split(header.AuthToken, "Bearer ")

	claims, err := authservice.ParseAuthToken(token[1])
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": err.Error(),
		})

		context.Abort()
		return
	}

	userId, err := uuid.Parse(claims.UserId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})

		context.Abort()
		return
	}

	user, err := userrepository.GetUserById(userId, db)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "O token de autenticação pertence a um usuário que não existe.",
		})

		context.Abort()
		return
	}

	context.Set("user", user)
}
