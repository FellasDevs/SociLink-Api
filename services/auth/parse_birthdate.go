package authservice

import "time"

func ParseBirthdate(birthdate string) (time.Time, error) {
	return time.Parse("2006-01-02", birthdate)
}
