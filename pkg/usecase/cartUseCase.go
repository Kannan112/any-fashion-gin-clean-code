package usecase

import (
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
func (c *CartUsecases) AddToCart(produtId, Userid int) error {
	err := c.cartRepo.AddToCart(produtId, Userid)
	return err
}
func (c *CartUsecases) RemoveFromCart(userId, productId int) error {
	err := c.cartRepo.RemoveFromCart(userId, productId)
	return err
}
