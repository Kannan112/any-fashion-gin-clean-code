package domain

type WishList struct {
	ID        uint `gorm:"primeryKey;not null"`
	ProductId uint
	Product   Product
	UsersID   uint
	Users     Users
}
