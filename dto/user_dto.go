package dto

type GetUserByIdDto struct {
	User PayloadUser
}

type GetUsersByNameDto struct {
	Users []PayloadUser
}
