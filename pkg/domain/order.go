package domain

import "time"

type OrderStatusType string

const ()

type Orders struct {
	Id          uint `gorm:"primaryKey;unique;not null"`
	UsersId     uint
	Users       Users
	OrderTime   time.Time
	AddressId   uint
	Address     Address
	CouponCode  string
	Coins       float64
	OrderTotal  float64
	OrderStatus string
}
type Order struct {
	Id          uint `gorm:"primaryKey;unique;not null"`
	UsersId     uint
	OrderTime   time.Time
	AddressId   uint
	OrderTotal  float32
	OrderStatus string
}
type OrderItem struct {
	Id            uint `gorm:"primaryKey;unique;not null"`
	OrdersId      uint
	Orders        Orders
	ProductItemId uint
	ProductItem   ProductItem
	Quantity      int
	Price         int
}

type OrderStatus struct {
	Id     uint `gorm:"primaryKey;unique;not null"`
	Status OrderStatusType
}
