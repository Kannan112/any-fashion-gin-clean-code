package interfaces

type OrderRepository interface{
	OrderAll(userId int)error
}