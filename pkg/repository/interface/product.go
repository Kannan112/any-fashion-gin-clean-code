package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

type ProductRepository interface {
	//Category
	CreateCategory(category req.Category) (res.Category, error)
	UpdateCategory(category req.Category, id int) (res.Category, error)
	DeleteCategory(id int) error
	DisplayCategory(id int) ([]res.Product, error)
	ListCategories(ctx context.Context, pagenation req.Pagenation) ([]res.Category, error)

	//Product
	AddProduct(product req.Product) (res.Product, error)
	UpdateProduct(id int, product req.Product) (res.Product, error)
	DeleteProduct(id int) error
	DeleteAllProducts() error
	DisplayProduct(id int) ([]res.Product, error)
	ListProducts() ([]res.Product, error)

	//Product-item
	AddProductItem(productItem req.ProductItem) (res.ProductItem, error)
	UpdateProductItem(id int, productItem req.ProductItem) (res.ProductItem, error)
	DeleteProductItem(id int) error
	DisaplyaAllProductItems(productId int) ([]domain.ProductItems, error)

	//Offer-side
	SaveOffer(ctx context.Context, offerdetails req.OfferTable) error
}
