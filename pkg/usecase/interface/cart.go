package interfaces

type CartUseCases interface {
	CreateCart(id int) error
	AddToCart(productId, userId int) error
	RemoveFromCart(productId, UserId int) error
	//ListCart(userId int)(res.ViewCart,error)
}
