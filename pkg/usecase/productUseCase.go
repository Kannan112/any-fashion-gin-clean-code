package usecase

import (
	"context"
	"errors"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type ProductUseCase struct {
	productRepo interfaces.ProductRepository
}

func NewProductUsecase(productRepo interfaces.ProductRepository) services.ProductUseCase {
	return &ProductUseCase{
		productRepo: productRepo,
	}
}

// CATEGORY
func (c *ProductUseCase) CreateCategory(category req.Category) (res.Category, error) {
	newCategort, err := c.productRepo.CreateCategory(category)
	if err != nil {
		return newCategort, errors.New("category already exists")
	}
	return newCategort, err
}
func (c *ProductUseCase) UpdateCategory(category req.Category, id int) (res.Category, error) {
	updateCategory, err := c.productRepo.UpdateCategory(category, id)
	return updateCategory, err
}
func (c *ProductUseCase) DeleteCategory(id int) error {
	err := c.productRepo.DeleteCategory(id)
	return err
}
func (c *ProductUseCase) ListCategories(ctx context.Context, pagenation req.Pagenation) ([]res.Category, error) {
	categories, err := c.productRepo.ListCategories(ctx, pagenation)
	return categories, err
}
func (c *ProductUseCase) DisplayCategory(id int) ([]res.Product, error) {
	var product []res.Product
	product, err := c.productRepo.DisplayCategory(id)
	return product, err
}

// PRODUCT
func (c ProductUseCase) AddProduct(product req.Product) (res.Product, error) {
	newProduct, err := c.productRepo.AddProduct(product)
	return newProduct, err
}
func (c ProductUseCase) UpdateProduct(id int, product req.Product) (res.Product, error) {
	updateProduct, err := c.productRepo.UpdateProduct(id, product)
	return updateProduct, err
}
func (c ProductUseCase) DeleteProduct(id int) error {
	err := c.productRepo.DeleteProduct(id)
	return err
}
func (c *ProductUseCase) DeleteAllProducts() error {
	err := c.productRepo.DeleteAllProducts()
	return err
}
func (c *ProductUseCase) ListProducts() ([]res.Product, error) {
	products, err := c.productRepo.ListProducts()
	return products, err
}
func (c *ProductUseCase) DisplayProduct(id int) ([]res.Product, error) {
	var product []res.Product
	product, err := c.productRepo.DisplayProduct(id)
	return product, err
}

// PRODUCT-ITEMS
func (c ProductUseCase) AddProductItem(productItem req.ProductItem) (res.ProductItem, error) {
	NewProductItem, err := c.productRepo.AddProductItem(productItem)
	return NewProductItem, err
}
func (c ProductUseCase) UpdateProductItem(productItem req.ProductItems) (res.ProductItem, error) {
	UpdateProductitem, err := c.productRepo.UpdateProductItem(productItem)
	return UpdateProductitem, err
}
func (c ProductUseCase) DeleteProductItem(id int) error {
	err := c.productRepo.DeleteProductItem(id)
	return err
}
func (c ProductUseCase) DisaplyaAllProductItems(productId int) ([]domain.ProductItems, error) {
	data, err := c.productRepo.DisaplyaAllProductItems(productId)
	return data, err
}
