package res

type CartItemsLim struct {
	ProductItemId uint `json:"product_item_id"`
	Quantity      int  `json:"quantity"`
}

type Display struct {
	ProductName string
	Brand       string
	Color       string
	Size        int
	Quantity    string
	Price       int
}

type ViewCart struct {
	CartItems []Display
	Subtotal  int
	Total     int
}
