package dto

import "SociLinkApi/models"

type GetMainTimelineResponseDto struct {
	Posts []PostResponseDto
}

type GetUserTimelineResponseDto struct {
	User  UserWithFriendsCountResponseDto
	Posts []PostResponseDto
}

type UserWithFriendsCount struct {
	models.User
	FriendsCount int
}

type UserWithFriendsCountResponseDto struct {
	UserResponseDto
	FriendsCount int
}

func UserToUserWithFriendsCountResponseDto(user UserWithFriendsCount) UserWithFriendsCountResponseDto {
	return UserWithFriendsCountResponseDto{
		UserResponseDto: UserToResponseDto(user.User),
		FriendsCount:    user.FriendsCount,
	}
}
