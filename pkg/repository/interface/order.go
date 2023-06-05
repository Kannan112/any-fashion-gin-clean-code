package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

type OrderRepository interface {
	OrderAll(userId int) (domain.Order, error)
	UserCancelOrder(orderId, userId int) (float32, error)
	ViewOrder(ctx context.Context) ([]domain.Order, error)
	ListAllOrders(userId int) ([]domain.Order, error)
	OrderDetails(ctx context.Context, orderId uint, userId uint) ([]res.UserOrder, error)
	ListOrderByCancelled(ctx context.Context) ([]domain.Order, error)
	ListOrderByPlaced(ctx context.Context) ([]domain.Order, error)
	//List Order by order delivered
	//ListOrderByDelivered(ctx context.Context) ([]domain.Order, ListOrderByCanceerror)
	
	//invoice download
	OrderInvoice(ctx context.Context,orderId uint)error

	//AdminCancelOrder(ctx context.Context, userId, orderId//ListOrders(ctx context.Context) ([]domain.Orders, error) int) error


	//List Order by order return
	//ListOrderByReturn(ctx context.Context) ([]domain.Order, error)
}
