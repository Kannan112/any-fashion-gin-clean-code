package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
)

type WalletUseCase interface {
	SaveWallet(ctx context.Context, userId int) error
	WallerProfile(ctx context.Context, userID uint) (res.Wallet, error)
	AddCoinToWallet(ctx context.Context, price float32, usersId uint) error
	ApplyWallet(ctx context.Context, userId uint) error
	RemoveWallet(ctx context.Context, userId uint) error
}
