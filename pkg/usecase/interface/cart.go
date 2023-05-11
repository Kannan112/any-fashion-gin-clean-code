package interfaces

import "github.com/kannan112/go-gin-clean-arch/pkg/domain"

type CartUseCases interface {
	CreateCart(id int) error
	AddToCart(productId, userId int) error
	RemoveFromCart(productId, UserId int) error
	ListCart(userId int) ([]domain.Cart, error)
}
