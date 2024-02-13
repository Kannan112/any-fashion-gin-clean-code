package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	handlerUtil "github.com/kannan112/go-gin-clean-arch/pkg/api/handlerUril"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
)

// RazorPayCheckout godoc
// @Summary Handle RazorPay checkout
// @Description Handle RazorPay checkout for a specific payment ID
// @Tags Payments
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Failure 401 {object} res.Response
// @Router /api/user/payment/razorpay-checkout [post]
func (cr *PaymentHandler) RazorPayCheckout(ctx *gin.Context) {
	//paramsId := ctx.Param("payment_id")
	UserId, err := handlerUtil.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 401,
			Message:    "Failed cookie token not found",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	razorPayOrder, err := cr.usecase.RazorPayCheckout(ctx, UserId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant place order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	fmt.Println("error not found")

	ctx.HTML(http.StatusOK, "razor.html", razorPayOrder)
}

// RazorPayVerify godoc
// @Summary Verify RazorPay payment
// @Description Verify RazorPay payment using the provided parameters
// @Tags Payments
// @Accept json
// @Produce json
// @Param razorpay_payment_id formData string true "RazorPay Payment ID"
// @Param razorpay_order_id formData string true "RazorPay Order ID"
// @Param razorpay_signature formData string true "RazorPay Signature"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /api/user/payment/razorpay-verify [post]
func (cr *PaymentHandler) RazorPayVerify(ctx *gin.Context) {
	razorPayPaymentId := ctx.Request.PostFormValue("razorpay_payment_id")
	razorPayOrderId := ctx.Request.PostFormValue("razorpay_order_id")
	razorpay_signature := ctx.Request.PostFormValue("razorpay_signature")

	userId, err := handlerUtil.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	body := req.RazorPayRequest{
		RazorPayPaymentId:  razorPayPaymentId,
		RazorPayOrderId:    razorPayOrderId,
		Razorpay_signature: razorpay_signature,
	}

	err = cr.usecase.VerifyRazorPay(ctx, body)
	if err != nil {
		fmt.Println("check One")
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    " faild to veify razorpay order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	order, err := cr.orderUsecase.OrderAll(userId)
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
