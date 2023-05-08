package repository

import (
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type OrderDatabase struct {
	DB gorm.DB
}

func NewOrderRepository(DB *gorm.DB) interfaces.OrderRepository {
	return &OrderDatabase{DB}
}

func (c *OrderDatabase) OrderAll(id int) {

}
