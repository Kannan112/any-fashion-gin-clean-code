package interfaces

type CartRepository interface {
	
	//Cart
	CreateCart(id int) error
	AddToCart(productId, userId int) error
	RemoveFromCart(produtId, userId int) error
	ListCart(userId int) error
}
