package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	handlerUtil "github.com/kannan112/go-gin-clean-arch/pkg/api/handlerUril"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type OrderHandler struct {
	orderUsecase  services.OrderUseCase
	walletUseCase services.WalletUseCase
}

func NewOrderHandler(orderUseCase services.OrderUseCase, walletUseCase services.WalletUseCase) *OrderHandler {
	return &OrderHandler{
		orderUsecase:  orderUseCase,
		walletUseCase: walletUseCase,
	}
}

// Cash on delevery
// @Summary order cart
// @ID order-all
// @Description order all cart items
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /user/cart/orderAll [get]
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

// User Cancel Order
// @Summary order cart
// @ID user-cancel-order
// @Description user cancel order
// @Tags Order
// @Accept json
// @Produce json
// @Param orderId path string true "orderId"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /user/order/cancel/{orderId} [patch]
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
	price, err := cr.orderUsecase.UserCancelOrder(orderId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't cancel order",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.walletUseCase.AddCoinToWallet(c, price, uint(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to add money",
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

// User List Order
// @Summary order status
// @ID list-all-order
// @Description user can view all his orders
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /user/order/listall [get]
func (ch *OrderHandler) ListAllOrders(c *gin.Context) {
	StartDateStr := c.Query("start")
	EndDateStr := c.Query("end")
	startDate, err := time.Parse("2006-7-5", StartDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to parse start date",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	endDate, err := time.Parse("2006-7-5", EndDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to parse end date",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userId, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "user login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	Details, err := ch.orderUsecase.ListAllOrders(userId, startDate, endDate)
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
func (cr *OrderHandler) OrderDetails(ctx *gin.Context) {
	paramsId := ctx.Param("orderId")
	orderId, err := strconv.Atoi(paramsId)
	fmt.Println(orderId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to find orderId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	UserId, err := handlerUtil.GetUserIdFromContext(ctx)
	//var data []res.OrderDetails
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to find userId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	data, err := cr.orderUsecase.OrderDetails(ctx, uint(orderId), uint(UserId))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to find",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "Order Details",
		Data:       data,
		Errors:     nil,
	})

}

// List Order By Placed
// @Summary order placed
// @ID list-order-by-placed
// @Description admin list order placed
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /admin/order/placed [get]
func (c *OrderHandler) ListOrderByPlaced(ctx *gin.Context) {
	data, err := c.orderUsecase.ListOrderByPlaced(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to collect data",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "order placed details",
		Data:       data,
		Errors:     nil,
	})

}

// List Order By Cancelled
// @Summary order cancelled
// @ID list-order-by-cancelled
// @Description admin view order cancelled
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /admin/order/cancelled [get]
func (c *OrderHandler) ListOrderByCancelled(ctx *gin.Context) {
	data, err := c.orderUsecase.ListOrderByCancelled(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to cancel",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "order canceled details",
		Data:       data,
		Errors:     nil,
	})
}

// Admin View Order
// @Summary order view
// @ID view-order
// @Description admin can view orders
// @Tags Order
// @Accept json
// @Produce json
// @Param start query string false "Start date (format: 2006-1-2)"
// @Param end query string false "End date (format: 2006-1-2)"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /admin/order [get]
func (c *OrderHandler) ViewOrder(ctx *gin.Context) {
	startDateStr := ctx.Query("start")
	endDateStr := ctx.Query("end")

	var startDate, endDate time.Time
	var err error

	if startDateStr != "" {
		startDate, err = time.Parse("2006-1-2", startDateStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, res.Response{
				StatusCode: 400,
				Message:    "failed to parse start date",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}

	if endDateStr != "" {
		endDate, err = time.Parse("2006-1-2", endDateStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, res.Response{
				StatusCode: 400,
				Message:    "failed to parse end date",
				Data:       nil,
				Errors:     err.Error(),
			})
			return
		}
	}
	viewOrder, err := c.orderUsecase.ViewOrder(ctx, startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to collect data",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "order details",
		Data:       viewOrder,
		Errors:     nil,
	})

}

func (c *OrderHandler) ListOrdersOfUsers(ctx *gin.Context) {
	userId, err := handlerUtil.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to get user id",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	orders, err := c.orderUsecase.ListOrdersOfUsers(ctx, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Failed to list",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "User order list",
		Data:       orders,
		Errors:     nil,
	})
}

// Admin Order Details
// @Summary order details
// @ID admin-order-details
// @Description admin can view orders
// @Tags Order
// @Accept json
// @Produce json
// @Param orderid path string true "orderid"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /admin/order/{orderid} [post]
func (c *OrderHandler) AdminOrderDetails(ctx *gin.Context) {
	orderId, err := strconv.Atoi(ctx.Param("orderid"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to get order id",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	order, err := c.orderUsecase.AdminOrderDetails(ctx, orderId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to get order details",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "order details",
		Data:       order,
		Errors:     nil,
	})
}
