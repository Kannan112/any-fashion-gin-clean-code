package handlerUtil

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserIdFromContext(c *gin.Context) (int, error) {
	id := c.Value("userId")
	userId, err := strconv.Atoi(fmt.Sprintf("%v", id))
	if err != nil {
		fmt.Println("1 test")
	}
	return userId, err
}
