package authcontroller

import (
	"SociLinkApi/dto"
	userrepository "SociLinkApi/repository/user"
	authservice "SociLinkApi/services/auth"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

func SignInController(userInfo dto.SignInDto, db *gorm.DB) (dto.SignInResponseDto, string, error) {
	user, err := userrepository.GetUserByEmail(userInfo.Email, db)
	if err != nil {
		return dto.SignInResponseDto{}, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInfo.Password)); err != nil {
		return dto.SignInResponseDto{}, "", err
	}

	token, err := authservice.CreateJWT(user.Id, time.Hour*24)

	signedUser := dto.SignInResponseDto{
		Name:      user.Name,
		Email:     user.Email,
		Birthdate: user.Birthdate.String(),
	}

	return signedUser, token, err
}
