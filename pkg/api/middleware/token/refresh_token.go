package token

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var refreshtokensrc = []byte("refresh-token-src")

func JWTRefreshTokenGen(userId int, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":   userId,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedstring, err := token.SignedString(refreshtokensrc)
	if err != nil {
		return "", err
	}
	return signedstring, nil
}
