package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
)

func (c *ProductHandler) SaveOffer(ctx *gin.Context) {
	// todo
	var body req.OfferTable
	err := ctx.Bind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to bind body",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	err = c.productuseCase.SaveOffer(ctx, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, res.Response{
			StatusCode: 400,
			Message:    "failed to save offer",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "success to save offer",
		Data:       nil,
		Errors:     nil,
	})
}
