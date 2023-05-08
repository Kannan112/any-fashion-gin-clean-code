package interfaces

type OrderUserCase interface{
	OrderAll(userId int)error
}