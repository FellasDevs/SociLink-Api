package authservice

import (
	authtypes "SociLinkApi/types/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"os"
	"time"
)

func CreateAuthToken(userId uuid.UUID, validFor time.Duration) (string, error) {
	jwtKey := []byte(os.Getenv("JWT_KEY"))

	claims := authtypes.CustomJWTClaims{
		UserId: userId.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(validFor)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}
