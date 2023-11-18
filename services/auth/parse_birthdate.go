package authservice

import (
	"errors"
	"time"
)

func ParseBirthdate(birthdate string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", birthdate)

	if err != nil {
		return time.Time{}, err
	}

	if date.After(time.Now()) {
		return time.Time{}, errors.New("data de nascimento não pode ser no futuro")
	}

	return date, nil
}
