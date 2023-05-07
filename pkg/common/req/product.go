package req

type Category struct {
	Name string `json:"name" validate:"required"`
}

type Product struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Brand       string `json:"brand" validate:"required"`
	CategoryId  string `json:"categoryid" validate:"required"`
}

type ProductItem struct {
	ID        uint
	ProductID uint
	Product   Product
	Model     string
	Size      int
	Material  string
	Gender    string
	SKU       string
	Color     string
	Qty       int
	Image     []string
	Price     int
}
