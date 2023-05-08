package res

type Display struct {
	ProductName string
	Brand string
	Color string
	Size int
	Quantity string
	Price int
}

type ViewCart struct {
	CartItems []Display
	Subtotal  int
	Total     int
}
