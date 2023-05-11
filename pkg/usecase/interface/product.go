package interfaces

import (
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

type ProductUseCase interface {
	CreateCategory(category req.Category) (res.Category, error)
	UpdateCategory(category req.Category, id int) (res.Category, error)
	DisplayCategory(id int) (res.Category, error)
	ListCategories() ([]res.Category, error)
	DeleteCategory(id int) error
	AddProduct(product req.Product) (res.Product, error)
	UpdateProduct(id int, product req.Product) (res.Product, error)
	DeleteProduct(id int) error
	AddProductItem(productItem req.ProductItem) (res.ProductItem, error)
	UpdateProductItem(id int, productItem req.ProductItem) (res.ProductItem, error)
	DeleteProductItem(id int) error
	DisaplyaAllProductItems(productId int) ([]domain.ProductItems, error)
}
