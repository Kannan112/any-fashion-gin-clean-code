package interfaces

import (
	"context"
	"time"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

type OrderUseCase interface {

	// userside
	OrderAll(id int) (domain.Order, error)
	UserCancelOrder(orderId, userId int) (float32, error)
	ListAllOrders(userId int, StartDate, EndDate time.Time) ([]domain.Order, error)
	RazorPayCheckout(ctx context.Context, userId int, paymentId int) (res.RazorPayResponse, error)
	VerifyRazorPay(ctx context.Context, body req.RazorPayRequest) error

	// common
	OrderDetails(ctx context.Context, OrderId uint, orderId uint) ([]res.UserOrder, error)

	// adminside
	ListOrderByPlaced(ctx context.Context) ([]domain.Order, error)
	ListOrderByCancelled(ctx context.Context) ([]domain.Order, error)
	ViewOrder(ctx context.Context, startDate, endDate time.Time) ([]domain.Order, error)
	AdminOrderDetails(ctx context.Context, orderId int) (res.OrderData, error)

	ListOrdersOfUsers(ctx context.Context, UserId int) ([]domain.Order, error)

}
