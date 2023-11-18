package usercontroller

import (
	"SociLinkApi/dto"
	userrepository "SociLinkApi/repository/user"
	authservice "SociLinkApi/services/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"net/url"
	"strings"
)

func EditUserInfo(context *gin.Context, db *gorm.DB) {
	uid, _ := context.Get("userId")

	var userInfo dto.EditUserInfoRequestDto

	if err := context.ShouldBindJSON(&userInfo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	user, err := userrepository.GetUserById(uid.(uuid.UUID), db)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	var fieldErrors []string

	if userInfo.Name != "" && userInfo.Name != user.Name {
		if len(userInfo.Name) < 6 {
			fieldErrors = append(fieldErrors, "Nome deve conter no mínimo 6 caracteres.")
		} else if len(userInfo.Name) > 50 {
			fieldErrors = append(fieldErrors, "Nome deve conter no máximo 50 caracteres.")
		} else {
			user.Name = userInfo.Name
		}
	}
	if userInfo.Nickname != "" && userInfo.Nickname != user.Nickname {
		if len(userInfo.Nickname) < 6 {
			fieldErrors = append(fieldErrors, "Nickname deve conter no mínimo 6 caracteres.")
		} else if len(userInfo.Nickname) > 50 {
			fieldErrors = append(fieldErrors, "Nickname deve conter no máximo 50 caracteres.")
		} else {
			user.Nickname = userInfo.Nickname
		}
	}
	if userInfo.Birthdate != "" {
		if birthdate, err := authservice.ParseBirthdate(userInfo.Birthdate); err != nil {
			fieldErrors = append(fieldErrors, "Data de nascimento inválida.")
		} else {
			user.Birthdate = birthdate
		}
	}
	if userInfo.Country != user.Country {
		if len(userInfo.Country) < 4 {
			fieldErrors = append(fieldErrors, "País deve conter no mínimo 4 caracteres.")
		} else if len(userInfo.Country) > 50 {
			fieldErrors = append(fieldErrors, "País deve conter no máximo 50 caracteres.")
		} else {
			user.Country = userInfo.Country
		}
	}
	if userInfo.City != user.City {
		if len(userInfo.City) < 4 {
			fieldErrors = append(fieldErrors, "Cidade deve conter no mínimo 4 caracteres.")
		} else if len(userInfo.City) > 50 {
			fieldErrors = append(fieldErrors, "Cidade deve conter no máximo 50 caracteres.")
		} else {
			user.City = userInfo.City
		}
	}
	if userInfo.Picture != user.Picture {
		if _, err := url.Parse(userInfo.Picture); err == nil || userInfo.Picture == "" {
			user.Picture = userInfo.Picture
		} else {
			fieldErrors = append(fieldErrors, "URL da foto de perfil inválida.")
		}
	}
	if userInfo.Banner != user.Banner {
		if _, err := url.Parse(userInfo.Banner); err == nil || userInfo.Banner == "" {
			user.Banner = userInfo.Banner
		} else {
			fieldErrors = append(fieldErrors, "URL do banner de perfil inválida.")
		}
	}

	if len(fieldErrors) > 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": strings.Join(fieldErrors, " "),
		})
		return
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
