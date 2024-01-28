package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	handlerUtil "github.com/kannan112/go-gin-clean-arch/pkg/api/handlerUril"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	"github.com/kannan112/go-gin-clean-arch/pkg/usecase"
)

// StripPaymentCheckout godoc
//
//	@Summary		Stripe checkout (User)
//	@Description	API for user to create stripe payment
//	@Tags			Payments
//	@Id				StripPaymentCheckout
//	@Param			shop_order_id	formData	string	true	"Shop Order ID"
//	@Router			/api/user/payment/stripe-checkout [post]
//	@Success		200	{object}	res.Response	"successfully stripe payment order created"
//	@Failure		500	{object}	res.Response	"Failed to create stripe order"
func (c *PaymentHandler) StripeCheckout(ctx *gin.Context) {
	id, err := handlerUtil.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to get user_id",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	stripeOrder, err := c.usecase.MakeStripeOrder(ctx, uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed complete payment",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	stripeResponse := res.OrderPayment{
		PaymentOrder: stripeOrder,
		PaymentType:  domain.StripePayment,
	}
	ctx.JSON(http.StatusOK, stripeResponse)

}

// StripePaymentVeify godoc
//
//	@Summary		Stripe verify (User)
//	@Description	API for user to callback backend after stripe payment for verification
//	@Tags			Payments
//	@Id				StripePaymentVeify
//	@Param			stripe_payment_id	formData	string	true	"Stripe payment ID"
//	@Router			/api/user/payment/stripe-verify [post]
//	@Success		200	{object}	res.Response	"Successfully stripe payment verified"
//	@Failure		402	{object}	res.Response	"Payment not approved"
//	@Failure		500	{object}	res.Response	"Failed to Approve order"
func (c *PaymentHandler) StripePaymentVerify(ctx *gin.Context) {
	stripePaymentID := ctx.Request.PostFormValue("stripe_payment_id")
	if stripePaymentID == "" {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to get from request body",
			Errors:     nil,
			Data:       nil,
		})
		return
	}

	userID, err := handlerUtil.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "failed to get user_id",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	err = c.usecase.VerifyStripeOrder(ctx, stripePaymentID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, usecase.ErrPaymentNotApproved) {
			statusCode = http.StatusPaymentRequired
		}
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: statusCode,
			Message:    "Failed to verify stripe payment",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	order, err := c.orderUsecase.OrderAll(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant place order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "orderplaced",
		Data:       order,
		Errors:     nil,
	})

}
