package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

type OrderUseCase interface {
	OrderAll(id int) (domain.Order, error)
	UserCancelOrder(orderId, userId int) error
	ListAllOrders(userId int) ([]domain.Orders, error)
	RazorPayCheckout(ctx context.Context, userId int, paymentId int) (res.RazorPayResponse, error)
	VerifyRazorPay(ctx context.Context, body req.RazorPayRequest) error
}
