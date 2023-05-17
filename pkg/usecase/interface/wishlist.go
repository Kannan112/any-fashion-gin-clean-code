package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
)

// continuee
type WishlistUseCases interface {
	AddToWishlist(itemId, userId int) error
	RemoveFromWishlist(ctx context.Context, userid, itemId int) error
	ListAllWishlist(ctx context.Context, userId int) ([]res.ProductItem, error)
}
