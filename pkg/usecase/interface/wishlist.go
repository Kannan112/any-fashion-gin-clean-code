package interfaces

// continuee
type WishlistUseCases interface {
	AddToWishlist(productId, userId int) error
	//(userId, productId int) error
	//	ViewFavourite(usersId)([]Response.)
}
