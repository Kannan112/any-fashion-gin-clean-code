package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
)

type WalletRepo interface {
	SaveWallet(ctx context.Context, userID int) error
	AddCoinToWallet(ctx context.Context, price float32, usersId uint) error
	WallerProfile(ctx context.Context, userID uint) (res.Wallet, error)
	ApplyWallet(ctx context.Context, userId uint) error
}
