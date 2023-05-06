package domain

type Carts struct {
	Id      uint `gorm:"primeryKey;unique;not null"`
	UsersID uint
	Users   Users
	Total   int
}
type CartItem struct {
	Id        uint `gorm:"primeryKey;unique;not null"`
	CartsID   uint
	Carts     Carts
	ProductID uint
	Product   Product
	Quantity  int
}
