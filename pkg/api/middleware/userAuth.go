package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserAuth(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	fmt.Println(authorizationHeader)
	if authorizationHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
		c.Abort()
		return
	}
	tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")

	userID, role, err := ValidateJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized%v", "err": err.Error()})
		c.Abort()
		return
	}
	if role != "user" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "role is not user%v", "err": err.Error()})
		c.Abort()
		return
	}
	c.Set("userId", userID)
	c.Next()
}
