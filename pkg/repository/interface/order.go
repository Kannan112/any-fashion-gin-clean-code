package interfaces

import (
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

type OrderRepository interface {
	OrderAll(userId int) (domain.Order, error)
	UserCancelOrder(orderId, userId int) (float32, error)
	ListAllOrders(userId int) ([]domain.Orders, error)

	//AdminCancelOrder(ctx context.Context, userId, orderId//ListOrders(ctx context.Context) ([]domain.Orders, error) int) error
}
