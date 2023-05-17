package interfaces

import "context"

// continuee
type WishlistUseCases interface {
	AddToWishlist(itemId, userId int) error
	RemoveFromWishlist(ctx context.Context, userid, itemId int) error
	//	ViewFavourite(usersId)([]Response.)
}
