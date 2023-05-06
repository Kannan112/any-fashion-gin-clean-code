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
	var exits bool
	query1 := `select exists(select 1 FROM categories where id=$1)`
	c.DB.Raw(query1, product.CategoryId).Scan(&exits)
	if !exits {
		return res.Product{}, fmt.Errorf("no category")
	}
	query := `INSERT INTO products (product_name,description,brand,qty,price,category_id,created_at)
		VALUES ($1,$2,$3,$4,$5,$6,NOW())
		RETURNING id,product_name AS name,description,brand,category_id`
	err := c.DB.Raw(query, product.Name, product.Description, product.Brand, product.Qty, product.Price, product.CategoryId).
		Scan(&newProduct).Error
	return newProduct, err
}
func (c *ProductDataBase) UpdateProduct(id int, product req.Product) (res.Product, error) {
	var updateProduct res.Product
	query := `UPDATE products SET product_name=$1,description=$2,brand=$3,qty=$4,price=$5,category_id=$6,updated_at=NOW()WHERE id=$5
	RETURNING id,product_name,description,brand,category_id`
	err := c.DB.Raw(query, product.Name, product.Description, product.Brand, product.Qty, product.Price, product.CategoryId).Scan(&updateProduct).Error
	return updateProduct, err
}
func (c *ProductDataBase) DeleteProduct(id int) error {
	query := `DELETE * FROM product WHERE id=$1`
	err := c.DB.Exec(query, id).Error
	return err
}
