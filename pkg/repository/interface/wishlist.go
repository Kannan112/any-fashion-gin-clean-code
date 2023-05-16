package interfaces

import "context"

type WishListRepo interface {
	AddToWishlist(ctx context.Context, id, productId int) error
	// RemoveFromWishlist(ctx context.Context, id int) error
}
