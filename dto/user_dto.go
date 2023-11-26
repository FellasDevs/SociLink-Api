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

type UserWithFriendsResponseDto struct {
	Id        string
	Name      string
	Nickname  string
	Birthdate time.Time
	Country   string
	City      string
	Picture   string
	Banner    string
	CreatedAt time.Time
	Friends   []UserResponseDto
}

type GetSelfResponseDto struct {
	User UserWithFriendsResponseDto
}

type GetUserByNicknameResponseDto struct {
	User UserWithFriendsResponseDto
}

type SearchUsersRequestDto struct {
	PaginationRequestDto
	Search string `form:"search"`
}

type SearchUsersResponseDto struct {
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

func UserToUserWithFriendsResponseDto(user models.User) UserWithFriendsResponseDto {
	friends := make([]UserResponseDto, len(user.Friends))

	for i, friend := range user.Friends {
		friends[i] = UserToUserResponseDto(*friend)
	}

	return UserWithFriendsResponseDto{
		Id:        user.ID.String(),
		Name:      user.Name,
		Nickname:  user.Nickname,
		Birthdate: user.Birthdate,
		Country:   user.Country,
		City:      user.City,
		Picture:   user.Picture,
		Banner:    user.Banner,
		CreatedAt: user.CreatedAt,
		Friends:   friends,
	}
}
