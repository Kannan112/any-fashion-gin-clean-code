package req

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFormValuesAsUint(ctx *gin.Context, name string) (uint, error) {

	value := ctx.Request.PostFormValue(name)
	uintVal, err := strconv.ParseUint(value, 10, 32)

	if err != nil || uintVal == 0 {
		return 0, fmt.Errorf("failed to get %s from request body as int", name)
	}

	return uint(uintVal), nil
}
