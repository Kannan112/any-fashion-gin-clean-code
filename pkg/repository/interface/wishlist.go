package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
)

type WishListRepo interface {
	AddToWishlist(id, productId int) error
	RemoveFromWishlist(ctx context.Context, userid, productid int) error
	ListAllWishlist(ctx context.Context, userId int, pagenation req.Pagenation) ([]res.ProductItem, error)
}
