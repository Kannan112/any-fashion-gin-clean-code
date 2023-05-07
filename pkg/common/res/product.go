package res

type Category struct {
	Id           int
	CategoryName string
}

type Product struct {
	Id           int `json:",omitempty"`
	Name         string
	Description  string
	Brand        string
	CategoryName string
}

type ProductItem struct {
	ID         uint
	ProductID  uint
	Product    Product
	Model      string
	Size       string
	Material   string
	Gender     string
	Type       string
	SKU        string
	QtyInStock int
	Price      int
}
