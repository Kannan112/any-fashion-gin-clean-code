package interfaces

import (
	"context"
)

type WishListRepo interface {
	AddToWishlist(id, productId int) error
	RemoveFromWishlist(ctx context.Context, userid, productid int) error
	//ListAllWishlist(ctx context.Context, userId int) error
}
