package domain

type Carts struct {
	Id        uint `gorm:"primeryKey;unique;not null"`
	UsersID   uint
	Users     Users
	CouponId  uint
	Sub_total int
	Coin      float32
	Total     int
}
type Cart struct {
	Id        uint `gorm:"primeryKey;unique;not null"`
	UsersID   uint
	Sub_total int
	Coin      float32
	Total     int
}
type CartItem struct {
	Id            uint `gorm:"primeryKey;unique;not null"`
	CartsID       uint
	Carts         Carts
	ProductItemID uint
	ProductItem   ProductItem
	Quantity      int
}
