package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type RenewHandler struct {
	TokenUseCase services.RenewTokenUseCase
}

func NewRenewHandler(token services.RenewTokenUseCase) *RenewHandler {
	return &RenewHandler{
		TokenUseCase: token,
	}
}

func (c *RenewHandler) GetAccessToken(ctx *gin.Context) {
	var Token req.AccessToken
	if err := gin.Bind(&Token); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, res.Response{
			StatusCode: 400,
			Message:    "binding failures",
			Data:       nil,
			Errors:     err,
		})
		return
	}

	NewAccessToken, err := c.TokenUseCase.GetAccessToken(ctx, Token.TokenString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to create a new access token",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "created new access token successfully",
		Data:       NewAccessToken,
		Errors:     nil,
	})
}
