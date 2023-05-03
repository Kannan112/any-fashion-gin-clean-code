package interfaces

type CartUseCases interface {
	CreateCart(id int) error
	AddToCart(ProductId, UserId int) error
	RemoveFromCart(productId, UserId int) error
}
