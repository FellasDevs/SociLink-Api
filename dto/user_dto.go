package dto

import (
	"SociLinkApi/models"
	"time"
)

type UserResponseDto struct {
	Id        string
	Name      string
	Nickname  string
	Birthdate time.Time
	Country   string
	City      string
	Picture   string
	Banner    string
	CreatedAt time.Time
}

type GetSelfResponseDto struct {
	User UserResponseDto
}

type GetUserByIdResponseDto struct {
	User UserResponseDto
}

type GetUsersByNameResponseDto struct {
	Users []UserResponseDto
}

type EditUserInfoRequestDto struct {
	Name      string
	Nickname  string
	Birthdate string
	Country   string
	City      string
	Picture   string
	Banner    string
}

type EditUserInfoResponseDto struct {
	User UserResponseDto
}

func UserToUserResponseDto(user models.User) UserResponseDto {
	return UserResponseDto{
		Id:        user.ID.String(),
		Name:      user.Name,
		Nickname:  user.Nickname,
		Birthdate: user.Birthdate,
		Country:   user.Country,
		City:      user.City,
		Picture:   user.Picture,
		Banner:    user.Banner,
		CreatedAt: user.CreatedAt,
	}
}
