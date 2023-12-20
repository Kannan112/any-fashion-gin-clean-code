package token

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// func GenerateToken(id int, exp time.Duration, role string) (string, error) {
func GenerateAccessToken(userID int, role string) (string, error) {
	// Create a new JWT claims instance
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 15).Unix(), // Token expiration time (24 hours)
		"iat":     time.Now().Unix(),                     // Token issuance time
	}

	// Create the JWT token with the claims and sign it with the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("access-token-src"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateRefreshToken(userID int, role string) (string, error) {
	// Create a new JWT claims instance
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 240).Unix(), // Token expiration time (24 hours)
		"iat":     time.Now().Unix(),                      // Token issuance time
	}

	// Create the JWT token with the claims and sign it with the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("refresh-token-src"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
