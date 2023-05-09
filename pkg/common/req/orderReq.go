package req

type Cart struct {
	Id     int
	Tottal int
}

type CartItems struct {
	ProductItemId int
	Quantity      int
	Price         int
	QntyInStock   int
}
