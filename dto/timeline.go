package dto

type GetMainTimelineResponseDto struct {
	Posts []PostResponseDto
}

type GetUserTimelineResponseDto struct {
	User  UserWithFriendsResponseDto
	Posts []PostResponseDto
}
