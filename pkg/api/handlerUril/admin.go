package handlerUtil

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAdminIdFromContext(c *gin.Context) (int, error) {
	id := c.Value("adminId")
	adminID, err := strconv.Atoi(fmt.Sprintf("%v", id))
	fmt.Println(c.Value("adminId"))
	if err != nil {
		fmt.Println(adminID)
	}
	return adminID, err
}
