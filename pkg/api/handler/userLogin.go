package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *UserHandler) LLLogin(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}
