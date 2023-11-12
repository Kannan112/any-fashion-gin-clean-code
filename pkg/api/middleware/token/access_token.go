package token

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var accesstokensrc = []byte("access-token-src")

func JWTAccessTokenGen(userId int, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":   userId,
		"role": role,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedstring, err := token.SignedString(accesstokensrc)
	if err != nil {
		return "", err
	}
	return signedstring, nil
}
