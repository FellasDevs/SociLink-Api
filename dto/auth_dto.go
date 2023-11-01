package dto

type SignUpDto struct {
	Email     string
	Name      string
	Birthdate string
	Password  string
}

type SignInDto struct {
	Email    string
	Password string
}
