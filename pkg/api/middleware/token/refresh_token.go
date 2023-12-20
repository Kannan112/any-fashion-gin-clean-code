package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
)

func RefreshTokenClaims(tokenString string) (res.TokenCalim, error) {
	var result res.TokenCalim

	Tokenvalue, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// validate the signing algorithm
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		// return the key for verification
		return []byte("refresh-token-src"), nil
	})
	if err != nil {
		return result, err
	}
	// check if the token is valid
	var parsedID interface{}
	var userrole interface{}

	if claims, ok := Tokenvalue.Claims.(jwt.MapClaims); ok && Tokenvalue.Valid {
		parsedID = claims["user_id"]
		userrole = claims["role"]
		//Check the expir
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return result, fmt.Errorf("your token expired please login again")
		}
		// fmt.Println(claims["exp"])
	}
	value, ok := parsedID.(float64)
	if !ok {
		return result, fmt.Errorf("expected an int value, but got %T", parsedID)
	}
	role, ok := userrole.(string)
	if !ok {
		return result, fmt.Errorf("expected an string value, but got %T", parsedID)
	}
	roleStr := string(role)
	idInt := int(value)

	result.ID = uint(idInt)
	result.Role = roleStr

	return result, err
}
