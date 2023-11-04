package authcontroller

import (
	"SociLinkApi/dto"
	"SociLinkApi/models"
	userrepository "SociLinkApi/repository/user"
	authservice "SociLinkApi/services/auth"
	"gorm.io/gorm"
	"time"
)

func SignUpController(userInfo dto.SignUpDto, db *gorm.DB) (models.User, string, error) {
	password, err := authservice.EncryptPassword(userInfo.Password)
	if err != nil {
		return models.User{}, "", err
	}

	birthdate, err := authservice.ParseBirthdate(userInfo.Birthdate)
	if err != nil {
		return models.User{}, "", err
	}

	user := models.User{
		Name:      userInfo.Name,
		Email:     userInfo.Email,
		Password:  password,
		Birthdate: birthdate,
	}

	createdUser, err := userrepository.CreateUser(user, db)
	if err != nil {
		return models.User{}, "", err
	}

	token, err := authservice.CreateJWT(createdUser.Id, time.Hour*24)

	return createdUser, token, err
}
