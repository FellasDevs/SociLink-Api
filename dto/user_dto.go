package dto

type GetUserByIdResponseDto struct {
	User PayloadUser
}

type GetUsersByNameResponseDto struct {
	Users []PayloadUser
}
