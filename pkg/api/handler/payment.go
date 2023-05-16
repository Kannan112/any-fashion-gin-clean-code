package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type PaymentHandler struct {
	usecase services.PaymentUsecases
}

func NewPaymentHandler(usercase services.PaymentUsecases) *PaymentHandler {
	return &PaymentHandler{usecase: usercase}
}

func (cr *PaymentHandler) SavePaymentMethod(c *gin.Context) {
	var paymentReq req.PaymentReq
	err := c.Bind(&paymentReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Cant bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.usecase.SavePaymentMethod(paymentReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to create payment method",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "successfully created",
		Data:       nil,
		Errors:     nil,
	})

}

func (cr *PaymentHandler) UpdatePaymentMethod(c *gin.Context) {
	var paymentReq req.PaymentReq
	err := c.Bind(&paymentReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Cant bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "pass a valide id",
			Data:       nil,
			Errors:     err,
		})
	}

	err = cr.usecase.UpdatePaymentMethod(id, paymentReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to update the payment method",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "successfully updated",
		Data:       nil,
		Errors:     nil,
	})
}
