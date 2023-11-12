package middleware

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func ValidateToken(tokenString string) (int, error) {
	Tokenvalue, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// validate the signing algorithm
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		// return the key for verification
		return []byte("access-token-src"), nil
	})
	if err != nil {
		return 0, err
	}
	// check if the token is valid
	var parsedID interface{}
	if claims, ok := Tokenvalue.Claims.(jwt.MapClaims); ok && Tokenvalue.Valid {
		parsedID = claims["id"]
		//Check the expir
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return 0, fmt.Errorf("access token expired please renew it")
		}
		// fmt.Println(claims["exp"])
	}
	value, ok := parsedID.(float64)
	if !ok {
		return 0, fmt.Errorf("expected an int value, but got %T", parsedID)
	}
	id := int(value)
	return id, err
}
