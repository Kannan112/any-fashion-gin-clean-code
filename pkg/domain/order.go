package domain

import "time"

type Orders struct {
	Id         uint `gorm:"primaryKey;unique;not null"`
	UsersId    uint
	Users      Users
	OrderTime  time.Time
	AddressId  uint
	Address    Address
	OrderTotal int
}
type Order struct {
	Id         uint `gorm:"primaryKey;unique;not null"`
	UsersId    uint
	OrderTime  time.Time
	AddressId  uint
	OrderTotal int
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
	Id     uint `gorma:"primaryKey;unique;not null"`
	Status string
}
