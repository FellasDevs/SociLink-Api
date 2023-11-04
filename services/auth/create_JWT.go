package authservice

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"time"
)

type CustomJWTClaims struct {
	UserId uuid.UUID
	jwt.RegisteredClaims
}

func CreateJWT(userId uuid.UUID, validFor time.Duration) (string, error) {
	jwtKey := []byte(os.Getenv("JWT_KEY"))

	claims := CustomJWTClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(validFor)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}
