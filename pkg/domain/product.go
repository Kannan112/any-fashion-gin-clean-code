package domain

import "time"

type Category struct {
	Id           uint   `gorm:"primaryKey;unique;not null"`
	Name string `gorm:"unique;not null"`
	Created_at   time.Time
	Updated_at   time.Time
}

type Product struct {
	Id          uint   `gorm:"primaryKey;unique;not null"`
	ProductName string `gorm:"unique;not null"`
	Description string
	Brand       string
	Qty         int
	Price       int
	Category_id uint
	Category    Category `gorm:"foreignKey:Category_id"`
	Created_at  time.Time
	Updated_at  time.Time
}

type Images struct {
	Id        uint `gorm:"primaryKey;unique;not null"`
	ProductId uint
	Product   Product `gorm:"foreignKey:ProductId"`
	FileName  string
}
