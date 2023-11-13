package dto

type UserResponseDto struct {
	Id        string
	Name      string
	Nickname  string
	Birthdate string
	Country   string
	City      string
	Picture   string
	Banner    string
	CreatedAt string
}

type GetUserByIdResponseDto struct {
	User UserResponseDto
}

type GetUsersByNameResponseDto struct {
	Users []UserResponseDto
}

type EditUserInfoRequestDto struct {
	Id        string
	Name      string
	Nickname  string
	Birthdate string
	Country   string
	City      string
	Picture   string
	Banner    string
}
