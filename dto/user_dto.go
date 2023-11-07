package dto

type UserResponseDto struct {
	Name      string
	Birthdate string
}

type GetUserByIdResponseDto struct {
	User UserResponseDto
}

type GetUsersByNameResponseDto struct {
	Users []UserResponseDto
}
