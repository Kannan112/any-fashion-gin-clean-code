package repository

import (
	"fmt"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type ProductDataBase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &ProductDataBase{DB}
}

func (c *ProductDataBase) CreateCategory(category req.Category) (res.Category, error) {
	var newCategoery res.Category
	query := `INSERT INTO categories(name,created_at)Values($1,NOW())RETURNING id,name`
	err := c.DB.Raw(query, category.Name).Scan(&newCategoery).Error
	return newCategoery, err
}
func (c *ProductDataBase) UpdateCategory(category req.Category, id int) (res.Category, error) {
	var updateCategory res.Category
	query := `UPDATE categories SET name=$1 WHERE id=$2 RETURNING id,name`
	err := c.DB.Raw(query, category.Name, id).Scan(&updateCategory).Error
	return updateCategory, err
}
func (c *ProductDataBase) DeleteCategory(id int) error {
	query := `DELETE FROM categories WHERE id=$1`
	err := c.DB.Exec(query, id).Error
	return err
}
func (c *ProductDataBase) ListCategories() ([]res.Category, error) {
	var categories []res.Category
	query := `SELECT * FROM categories`
	err := c.DB.Raw(query).Scan(&categories).Error
	return categories, err
}
func (c *ProductDataBase) DisplayCategory(id int) (res.Category, error) {
	var category res.Category
	query := `SELECT * FROM categories WHERE id=$1`
	err := c.DB.Raw(query, id).Scan(&category).Error
	return category, err
}
func (c *ProductDataBase) AddProduct(product req.Product) (res.Product, error) {
	var newProduct res.Product
	var exists bool

	// Check if category exists
	query1 := `SELECT EXISTS(SELECT 1 FROM categories WHERE id=?)`
	c.DB.Raw(query1, product.CategoryId).Scan(&exists)
	if !exists {
		return res.Product{}, fmt.Errorf("category does not exist")
	}

	// Insert new product
	query := `INSERT INTO products (product_name, description, brand, category_id, created_at)
			  VALUES ($1, $2, $3, $4, NOW())
			  RETURNING id, product_name AS name, description, brand, category_id`
	err := c.DB.Raw(query, product.Name, product.Description, product.Brand, product.CategoryId).
		Scan(&newProduct).Error
	return newProduct, err
}
func (c *ProductDataBase) UpdateProduct(id int, product req.Product) (res.Product, error) {
	var updatedProduct res.Product
	query := `
		UPDATE products
		SET product_name = $1,
			description = $2,
			brand = $3,
			category_id = $4,
			updated_at = NOW()
		WHERE id = $5
		RETURNING id, product_name, description, brand, category_id
	`
	err := c.DB.Raw(query, product.Name, product.Description, product.Brand, product.CategoryId, id).
		Scan(&updatedProduct).Error
	return updatedProduct, err
}
func (c *ProductDataBase) DeleteProduct(id int) error {
	query := `DELETE FROM product WHERE id=$1`
	err := c.DB.Exec(query, id).Error
	return err
}
func (c *ProductDataBase) AddProductItem(productItem req.ProductItem) (res.ProductItem, error) {
	var NewProductItem res.ProductItem
	query := `INSERT INTO product_items SET product_id=$1,sku=$2,qnty_in_stock=$3,gender=$4,model=$5,size=$6,color=$7,material=$8,price$9,update_at=NOW())
	RETURNING product_id,qnty_in_stock,gender,model,size,color,material,price,created_at`
	err := c.DB.Raw(query, productItem.ProductID, productItem.SKU, productItem.Qty, productItem.Gender, productItem.Model, productItem.Size, productItem.Color, productItem.Material, productItem.Price).Scan(&NewProductItem).Error
	return NewProductItem, err
}
func (c *ProductDataBase) UpdateProductItem(id int, productItem req.ProductItem) (res.ProductItem, error) {
	var UpdateProductItem res.ProductItem
	query := `UPDATE product_items(product_id,sku,qnty_in_stock,gender,model,size,color,material,price,created_at)SET($1,$2,$3,$4,$5,$6,$7,$8,$9,NOW())
	RETURNING product_id,qnty_in_stock,gender,model,size,color,material,price,created_at`
	err := c.DB.Raw(query, productItem.ProductID, productItem.SKU, productItem.Qty, productItem.Gender, productItem.Model, productItem.Size, productItem.Color, productItem.Material, productItem.Price).Scan(&UpdateProductItem).Error
	return UpdateProductItem, err
}
func (c *ProductDataBase) DeleteProductItem(id int) error {
	delete := `DELETE FROM product_items WHERE id=$1`
	err := c.DB.Exec(delete, id).Error
	return err
}
