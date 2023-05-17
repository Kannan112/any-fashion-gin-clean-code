package domain

type WishList struct {
	ID          uint `gorm:"primeryKey;not null"`
	UsersId     uint
	Users       Users `gorm:"foreignKey:UsersId"`
	ItemId      uint
	ProductItem ProductItem `gorm:"foreignKey:ItemId"`
}
