package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserAuth(c *gin.Context) {
	tokenString, err := c.Cookie("UserAuth")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	userId, err := ValidateToken(tokenString)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set("userId", userId)
	fmt.Println("user_id", userId)
	c.Next()

}
