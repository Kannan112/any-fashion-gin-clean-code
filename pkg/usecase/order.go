package usecase

import (
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type OrderUseCase struct {
	orderRepo interfaces.OrderRepository
}

func NewOrderUseCase(orderRepo interfaces.OrderRepository) services.OrderUseCase {
	return &OrderUseCase{
		orderRepo: orderRepo,
	}
}
func (c *OrderUseCase) OrderAll(id int) (domain.Orders, error) {
	order, err := c.orderRepo.OrderAll(id)
	return order, err
}
func (c *OrderUseCase) UserCancelOrder(orderId, userId int) error {
	err := c.orderRepo.UserCancelOrder(orderId, userId)
	return err
}
func (c *OrderUseCase) ListAllOrders(userId int) ([]domain.Order, error) {
	order, err := c.orderRepo.ListAllOrders(userId)
	return order, err
}
