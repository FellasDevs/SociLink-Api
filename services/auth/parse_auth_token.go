package authservice

import (
	authtypes "SociLinkApi/types/auth"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func ParseAuthToken(jwtString string) (*authtypes.CustomJWTClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	}

	token, err := jwt.ParseWithClaims(jwtString, &authtypes.CustomJWTClaims{}, keyFunc)

	claims, ok := token.Claims.(*authtypes.CustomJWTClaims)

	if err != nil {
		return claims, err
	} else if !ok || !token.Valid {
		return claims, errors.New("token inválido")
	}

	return claims, nil
}
