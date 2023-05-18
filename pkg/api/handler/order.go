package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	handlerUtil "github.com/kannan112/go-gin-clean-arch/pkg/api/handlerUril"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type OrderHandler struct {
	orderUsecase services.OrderUseCase
}

func NewOrderHandler(orderUseCase services.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUsecase: orderUseCase,
	}
}
func (cr *OrderHandler) OrderAll(c *gin.Context) {

	userId, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	order, err := cr.orderUsecase.OrderAll(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant place order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "orderplaced",
		Data:       order,
		Errors:     nil,
	})
}

func (cr *OrderHandler) UserCancelOrder(c *gin.Context) {
	userId, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := c.Param("orderId")
	orderId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.orderUsecase.UserCancelOrder(orderId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't cancel order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "order canceld",
		Data:       nil,
		Errors:     nil,
	})
}
func (ch *OrderHandler) ListAllOrders(c *gin.Context) {
	userId, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	Details, err := ch.orderUsecase.ListAllOrders(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to list",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "List All Orders",
		Data:       Details,
		Errors:     nil,
	})
}

func (cr *OrderHandler) RazorPayCheckout(ctx *gin.Context) {
	UserID, err := handlerUtil.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := ctx.Param("payment_id")
	payment_id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	razorPayOrder, err := cr.orderUsecase.RazorPayCheckout(ctx, UserID, payment_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant place order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.HTML(http.StatusOK, "razor.html", razorPayOrder)
}

func (cr *OrderHandler) RazorPayVerify(ctx *gin.Context) {
	razorPayPaymentId := ctx.Request.PostFormValue("razorpay_payment_id")
	razorPayOrderId := ctx.Request.PostFormValue("razorpay_order_id")
	razorpay_signature := ctx.Request.PostFormValue("razorpay_signature")
	// paramsId := ctx.Request.PostFormValue("payment_id")

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
	// paymentid, err := strconv.Atoi(paramsId)

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

	err = cr.orderUsecase.VerifyRazorPay(ctx, body)
	if err != nil {
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
