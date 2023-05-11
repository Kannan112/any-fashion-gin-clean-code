package interfaces

import "github.com/kannan112/go-gin-clean-arch/pkg/domain"

type OrderUseCase interface {
	OrderAll(id int) (domain.Orders, error)
	UserCancelOrder(orderId, userId int) error
	ListAllOrders(userId int) ([]domain.Order, error)
}
