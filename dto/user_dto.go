package dto

type UserResponseDto struct {
	Id        string
	Name      string
	Birthdate string
}

type GetUserByIdResponseDto struct {
	User UserResponseDto
}

type GetUsersByNameResponseDto struct {
	Users []UserResponseDto
}
