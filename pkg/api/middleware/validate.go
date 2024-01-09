package middleware

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func ValidateJWT(TokenString string) (int, string, error) {
	tokenValue, err := jwt.Parse(TokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Use your actual secret key here instead of "access"
		return []byte("access-token-src"), nil
	})

	if err != nil {
		return 0, "", fmt.Errorf("JWT validation failed: %v", err.Error())
	}

	if claims, ok := tokenValue.Claims.(jwt.MapClaims); ok && tokenValue.Valid {
		role, roleOk := claims["role"].(string)
		if !roleOk {
			return 0, "", errors.New("failed to get a role")
		}

		paramsId, idOK := claims["user_id"].(float64)

		exp, expOK := claims["exp"].(float64)
		if !idOK || !expOK {

			return 0, "", errors.New("missing or invalid claims in the jwt")
		}

		if float64(time.Now().Unix()) > exp {

			return 0, "", errors.New("token has expired")
		}

		return int(paramsId), role, nil
	}
	return 0, "", fmt.Errorf("JWT clai,ms are not valid")
}
