package usecase

import (
	"context"

	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type WishListUseCase struct {
	wishListRepo interfaces.WishListRepo
}

func NewWishlistUsecase(repo interfaces.WishListRepo) services.WishlistUseCases {
	return &WishListUseCase{
		wishListRepo: repo,
	}
}
func (c *WishListUseCase) AddToWishlist(productId, userId int) error {
	err := c.wishListRepo.AddToWishlist(userId, productId)
	return err
}
func (c *WishListUseCase) RemoveFromWishlist(ctx context.Context, userid, productid int) error {
	err := c.wishListRepo.RemoveFromWishlist(ctx, userid, productid)
	return err
}
