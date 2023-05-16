package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

type CartRepository interface {
	FindCart(ctx context.Context, userId int) (domain.Cart, error)
	CreateCart(id int) error
	AddToCart(productId, userId int) error
	RemoveFromCart(produtId, userId int) error
	ListCart(userId int) ([]domain.Cart, error)
}
