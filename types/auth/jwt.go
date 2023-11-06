package authtypes

import (
	"github.com/golang-jwt/jwt/v5"
)

type CustomJWTClaims struct {
	UserId string
	jwt.RegisteredClaims
}
