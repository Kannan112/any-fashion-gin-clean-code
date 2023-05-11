package interfaces

import "github.com/kannan112/go-gin-clean-arch/pkg/domain"

type CartRepository interface {
	CreateCart(id int) error
	AddToCart(productId, userId int) error
	RemoveFromCart(produtId, userId int) error
	ListCart(userId int) ([]domain.Cart, error)
}
