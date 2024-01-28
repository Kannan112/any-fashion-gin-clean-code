package usecase

import (
	"context"
	"time"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type OrderUseCase struct {
	orderRepo interfaces.OrderRepository
	cartRepo  interfaces.CartRepository
	userRepo  interfaces.UserRepository
}

func NewOrderUseCase(orderRepo interfaces.OrderRepository, cartRepo interfaces.CartRepository, userRepo interfaces.UserRepository) services.OrderUseCase {
	return &OrderUseCase{
		orderRepo: orderRepo,
		cartRepo:  cartRepo,
		userRepo:  userRepo,
	}
}

func (c *OrderUseCase) OrderAll(id int) (domain.Order, error) {
	data, err := c.orderRepo.OrderAll(id)
	return data, err
}

func (c *OrderUseCase) UserCancelOrder(orderId, userId int) (float32, error) {
	price, err := c.orderRepo.UserCancelOrder(orderId, userId)
	if err != nil {
		return 0, err
	}
	return price, err
}

func (c *OrderUseCase) ListAllOrders(userId int, startDate, EndDate time.Time) ([]domain.Order, error) {
	order, err := c.orderRepo.ListAllOrders(userId, startDate, EndDate)
	return order, err
}

func (c *OrderUseCase) OrderDetails(ctx context.Context, orderId uint, userId uint) ([]res.UserOrder, error) {
	var OrderDetails []res.UserOrder
	OrderDetails, err := c.orderRepo.OrderDetails(ctx, orderId, userId)
	if err != nil {
		return OrderDetails, err
	}
	return OrderDetails, err

}

func (c *OrderUseCase) ListOrderByPlaced(ctx context.Context) ([]domain.Order, error) {
	var OrderDetails []domain.Order
	OrderDetails, err := c.orderRepo.ListOrderByPlaced(ctx)
	if err != nil {
		return OrderDetails, err
	}
	return OrderDetails, err
}

func (c *OrderUseCase) ListOrderByCancelled(ctx context.Context) ([]domain.Order, error) {
	var OrderDetails []domain.Order
	OrderDetails, err := c.orderRepo.ListOrderByCancelled(ctx)
	if err != nil {
		return OrderDetails, err
	}
	return OrderDetails, err
}

func (c *OrderUseCase) ViewOrder(ctx context.Context, startDate, endDate time.Time) ([]domain.Order, error) {
	var OrderDetails []domain.Order
	OrderDetails, err := c.orderRepo.ViewOrder(ctx, startDate, endDate)
	if err != nil {
		return OrderDetails, err
	}
	return OrderDetails, err
}

func (c *OrderUseCase) ListOrdersOfUsers(ctx context.Context, UserId int) ([]domain.Order, error) {
	order, err := c.orderRepo.ListOrdersOfUsers(ctx, UserId)
	return order, err
}

func (c *OrderUseCase) AdminOrderDetails(ctx context.Context, orderId int) (res.OrderData, error) {
	order, err := c.orderRepo.AdminOrderDetails(ctx, orderId)
	return order, err
}
