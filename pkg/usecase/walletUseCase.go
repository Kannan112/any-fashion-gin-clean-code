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

func (uc *WalletUseCase) SaveWallet(ctx context.Context, userID int) error {
	err := uc.WalletRepo.SaveWallet(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *WalletUseCase) GetWalletProfile(ctx context.Context, userID uint) (res.Wallet, error) {
	response, err := uc.WalletRepo.GetWalletProfile(ctx, userID)
	if err != nil {
		return response, err
	}
	return response, nil
}

func (uc *WalletUseCase) AddCoinToWallet(ctx context.Context, price float32, userID uint) error {
	err := uc.WalletRepo.AddCoinToWallet(ctx, price, userID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *WalletUseCase) ApplyWallet(ctx context.Context, userID uint) error {
	err := uc.WalletRepo.ApplyWallet(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}

func (uc *WalletUseCase) RemoveWallet(ctx context.Context, userID uint) error {
	err := uc.WalletRepo.RemoveWallet(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
