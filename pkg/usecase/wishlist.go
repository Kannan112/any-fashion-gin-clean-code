package usecase

import (
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type WishListUseCase struct {
	wishListUsecase interfaces.WishListRepo
}

func NewWishlistUsecase(repo interfaces.WishListRepo) services.WishlistUseCases {
	return &WishListUseCase{
		wishListUsecase: repo,
	}
}
func (c *WishListUseCase) AddToWishlist(productId, userId int) error {
	err := c.AddToWishlist(userId, productId)
	return err
}
