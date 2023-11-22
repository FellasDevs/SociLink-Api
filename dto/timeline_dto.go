package dto

import types "SociLinkApi/types/pagination"

type GetMainTimelineResponseDto struct {
	types.PaginationResponse
	Posts []PostResponseDto
}

type GetUserTimelineResponseDto struct {
	types.PaginationResponse
	User  UserWithFriendsResponseDto
	Posts []PostResponseDto
}
