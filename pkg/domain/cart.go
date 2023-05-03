package domain

type Carts struct {
	Id       uint `gorm:"primeryKey;unique;not null"`
	Users_id uint
	Users    Users `gorm:"foreignKey:User_id"`
	Total    int
}
type CartItem struct {
	Id         uint `gorm:"primeryKey;unique;not null"`
	Carts_id   uint
	Carts      Carts `gorm:"foreignKey:Carts_id"`
	Product_id uint
	Product    Product `gorm:"foreignKey:Product_id"`
	Quantity   int
}
