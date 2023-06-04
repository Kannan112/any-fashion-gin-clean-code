package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

type CartRepository interface {
	FindCart(ctx context.Context, userId int) (domain.Carts, error)
	CreateCart(id int) error
	AddToCart(productId, userId int) error
	RemoveFromCart(produtId, userId int) error
	ListCart(userId int) ([]domain.Cart, error)
	ListCartItems(ctx context.Context, userId int) ([]res.Display, error)
}
