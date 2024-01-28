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
	usecase      services.PaymentUsecases
	orderUsecase services.OrderUseCase
}

func NewPaymentHandler(usercase services.PaymentUsecases, orderusecase services.OrderUseCase) *PaymentHandler {
	return &PaymentHandler{
		usecase:      usercase,
		orderUsecase: orderusecase,
	}
}

// CartOrderPaymentSelectPage godoc
//
//		@Summary		Render Payment Page (User)
//	 @Security BearerTokenAuth
//		@Description	API for user to render payment select page
//		@Id				CartOrderPaymentSelectPage
//		@Tags			Payments
//		@Router			/api/user/payment/checkout/payment-select-page [get]
//		@Success		200	{object}	res.Response{}	"Successfully rendered payment page"
//		@Failure		500	{object}	res.Response{}	"Failed to render payment page"
func (c *PaymentHandler) CartOrderPaymentSelectPage(ctx *gin.Context) {

	Payments, err := c.usecase.ListPaymentMethod(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to render payment page",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.HTML(200, "paymentForm.html", Payments)
}

// GetAllPaymentMethodsUser godoc
//
//	@summary		Get payment methods
//	@Description	API for user to get all payment methods
//	@tags		    Payments
//	@id				GetAllPaymentMethodsUser
//	@Router			/api/user/payment/payment-methods [get]
//	@Success		200	{object}	res.Response{}	"Failed to retrieve payment methods"
//	@Failure		500	{object}	res.Response{}	"Successfully retrieved all payment methods"
func (c *PaymentHandler) GetPaymentMethodUser() func(ctx *gin.Context) {
	return c.ListPayment()
}

// GetAllPaymentMethodsUser godoc
//
//	@summary		Get payment methods
//	@Description	API for admin to get all payment methods
//	@tags		    Payments
//	@id				GetAllPaymentMethodsAdmin
//	@Router			/api/admin/payment-methods [get]
//
// @Security BearerTokenAuth
// @Success		200	{object}	res.Response{}	"Failed to retrieve payment methods"
// @Failure		500	{object}	res.Response{}	"Successfully retrieved all payment methods"
func (c *PaymentHandler) GetPaymentMethodAdmin() func(ctx *gin.Context) {
	return c.ListPayment()
}

func (c *PaymentHandler) ListPayment() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		result, err := c.usecase.ListPaymentMethod(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, res.Response{
				StatusCode: 400,
				Message:    "failed to list",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusAccepted, res.Response{
			StatusCode: 200,
			Message:    "list of payment methods",
			Data:       result,
			Errors:     nil,
		})
	}

}

// UpdatePaymentMethod godoc
// @Summary Update a payment method
// @Description Update a payment method by providing the payment ID and request payload
// @Tags Payments
// @Accept json
// @Produce json
// @Param id path int true "Payment ID" Format(int)
// @Param request body req.PaymentReq true "Payment request payload"
// @Success 202 {object} res.Response
// @Failure 400 {object} res.Response
// @Security BearerTokenAuth
// @Router /api/admin/payment-update/{id} [put]
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
