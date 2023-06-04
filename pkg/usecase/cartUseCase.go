package usecase

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type CartUsecases struct {
	cartRepo interfaces.CartRepository
}

func NewCartUseCase(cartRepo interfaces.CartRepository) services.CartUseCases {
	return &CartUsecases{
		cartRepo: cartRepo,
	}
}

func (c *CartUsecases) CreateCart(id int) error {
	err := c.cartRepo.CreateCart(id)
	return err
}
func (c *CartUsecases) AddToCart(produtId, userId int) error {
	err := c.cartRepo.AddToCart(produtId, userId)
	return err
}
func (c *CartUsecases) RemoveFromCart(userId, productId int) error {
	err := c.cartRepo.RemoveFromCart(userId, productId)
	return err
}
func (c *CartUsecases) ListCart(userId int) ([]domain.Cart, error) {
	list, err := c.cartRepo.ListCart(userId)
	return list, err

}
func (c *CartUsecases) ListCartItems(ctx context.Context, userId int) ([]res.Display, error) {
	data, err := c.cartRepo.ListCartItems(ctx, userId)
	return data, err
}
