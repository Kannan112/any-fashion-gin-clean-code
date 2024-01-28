package res

type CartItemsLim struct {
	ProductItemId uint `json:"product_item_id"`
	Quantity      int  `json:"quantity"`
}

type Display struct {
	ProductName string `json:"product_name"`
	Brand       string `json:"brand"`
	Color       string `json:"color"`
	Size        int    `json:"size"`
	Price       int    `json:"price"`
}

type ViewCart struct {
	CartItems []Display
	Subtotal  int
	Total     int
}
