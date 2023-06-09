package usecase

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type WalletUseCase struct {
	WalletRepo interfaces.WalletRepo
}

func NewWalletUseCase(repo interfaces.WalletRepo) services.WalletUseCase {
	return &WalletUseCase{
		WalletRepo: repo,
	}
}

func (c *WalletUseCase) SaveWallet(ctx context.Context, userId int) error {
	err := c.WalletRepo.SaveWallet(ctx, userId)
	return err
}
func (c *WalletUseCase) WallerProfile(ctx context.Context, userID uint) (res.Wallet, error) {
	response, err := c.WalletRepo.WallerProfile(ctx, userID)
	return response, err
}
func (c *WalletUseCase) AddCoinToWallet(ctx context.Context, price float32, usersId uint) error {
	err := c.WalletRepo.AddCoinToWallet(ctx, price, usersId)
	return err
}
func (c *WalletUseCase) ApplyWallet(ctx context.Context, userId uint) error {
	err := c.WalletRepo.ApplyWallet(ctx, userId)
	return err
}
func (c *WalletUseCase) RemoveWallet(ctx context.Context, userId uint) error {
	err := c.WalletRepo.RemoveWallet(ctx, userId)
	return err
}
