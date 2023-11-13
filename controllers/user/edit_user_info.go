package usercontroller

import (
	"SociLinkApi/dto"
	userrepository "SociLinkApi/repository/user"
	authservice "SociLinkApi/services/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func EditUserInfo(context *gin.Context, db *gorm.DB) {
	var userInfo dto.EditUserInfoRequestDto

	if err := context.ShouldBindJSON(&userInfo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	userId, err := uuid.Parse(userInfo.Id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	user, err := userrepository.GetUserById(userId, db)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	if userInfo.Name != "" {
		user.Name = userInfo.Name
	}
	if userInfo.Nickname != "" {
		user.Nickname = userInfo.Nickname
	}
	if userInfo.Birthdate != "" {
		birthdate, err := authservice.ParseBirthdate(userInfo.Birthdate)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}

		user.Birthdate = birthdate
	}
	if userInfo.Country != "" {
		user.Country = userInfo.Country
	}
	if userInfo.City != "" {
		user.City = userInfo.City
	}
	if userInfo.Picture != "" {
		user.Picture = userInfo.Picture
	}
	if userInfo.Banner != "" {
		user.Banner = userInfo.Banner
	}

	err = userrepository.UpdateUser(&user, db)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "User info updated successfully",
		})
		return
	}
}
