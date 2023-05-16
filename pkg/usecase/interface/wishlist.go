package interfaces

import "context"

// continuee
type WishlistUseCases interface {
	AddToWishlist(productId, userId int) error
	RemoveFromWishlist(ctx context.Context, userid, productid int) error
	//	ViewFavourite(usersId)([]Response.)
}
