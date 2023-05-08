package domain

import "time"

type Category struct {
	Id         uint   `gorm:"primaryKey;unique;not null"`
	Name       string `gorm:"unique;not null"`
	Created_at time.Time
	Updated_at time.Time
}

type Product struct {
	Id          uint   `gorm:"primaryKey;unique;not null"`
	ProductName string `gorm:"unique;not null"`
	Description string
	Brand       string
	CategoryID  uint
	Category    Category
	Created_at  time.Time
	Updated_at  time.Time
}

type ProductItem struct {
	ID          uint `gorm:"primaryKey;unique;not null" json:"id"`
	ProductID   uint `gorm:"not null" json:"product_id" validate:"required"`
	Product     Product
	SKU         string
	QntyInStock int     `gorm:"not null" json:"qnty_in_stock" validate:"required"`
	Gender      string  `gorm:"not null" json:"gender" validate:"required"`
	Model       string  `gorm:"not null" json:"model" validate:"required"`
	Size        int     `gorm:"not null" json:"size" validate:"required"`
	Color       string  `gorm:"not null" json:"color" validate:"required"`
	Material    string  `gorm:"not null" json:"material" validate:"required"`
	Price       float64 `gorm:"not null" json:"price" validate:"required"`
	Created_at  time.Time
}

type Images struct {
	Id            uint `gorm:"primaryKey;unique;not null"`
	ProductItemID uint `gorm:"not null" json:"product_item_id" validate:"required"`
	ProductItem   ProductItem
	FileName      string `json:"file_name"`
}
