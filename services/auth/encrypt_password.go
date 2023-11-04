package authservice

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(passwordString string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(passwordString), 10)

	return string(password), err
}
