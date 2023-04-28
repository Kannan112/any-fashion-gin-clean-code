package usecase

import (
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
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

func (c *ProductUseCase) CreateCategory(category req.Category) (res.Category, error) {
	newCategort, err := c.productRepo.CreateCategory(category)
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
func (c *ProductUseCase) ListCategories() ([]res.Category, error) {
	categories, err := c.productRepo.ListCategories()
	return categories, err
}
func (c *ProductUseCase) DisplayCategory(id int) (res.Category, error) {
	category, err := c.productRepo.DisplayCategory(id)
	return category, err

}
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
