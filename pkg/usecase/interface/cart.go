package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

type CartUseCases interface {
	CreateCart(id int) error
	AddToCart(productId, userId int) error
	RemoveFromCart(productId, UserId int) error
	ListCart(userId int) ([]domain.Cart, error)
	ListCartItems(ctx context.Context, userId int, pagenation req.Pagenation) ([]res.Display, error)
}
