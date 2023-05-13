package handler

import (
	"net/http"

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
	}
	err = cr.usecase.SavePaymentMethod(paymentReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to create payment method",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "successfully created",
		Data:       nil,
		Errors:     err.Error(),
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
	}
	err = cr.usecase.UpdatePaymentMethod(paymentReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to update the payment method",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "successfully updated",
		Data:       nil,
		Errors:     err.Error(),
	})
}
