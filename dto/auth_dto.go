package dto

type SignUpRequestDto struct {
	Email     string
	Name      string
	Birthdate string
	Password  string
}

type SignInRequestDto struct {
	Email    string
	Password string
}

type PayloadUser struct {
	Name      string
	Email     string
	Birthdate string
}

type SignInResponseDto struct {
	User      PayloadUser
	AuthToken string
}

type SignUpResponseDto struct {
	User      PayloadUser
	AuthToken string
}
