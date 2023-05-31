package domain

type Wallet struct {
	Id      uint
	UsersId uint
	Users   Users
	Coins   float32
}
