package interfaces

import "github.com/kannan112/go-gin-clean-arch/pkg/domain"

type OrderRepository interface {
	OrderAll(id int) (domain.Orders, error)
	UserCancelOrder(orderId, userId int) error
	ListAllOrders(userId int) ([]domain.Orders, error)

}
