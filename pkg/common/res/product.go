package res

type Category struct {
	Id   int
	Name string
}

type Product struct {
	Id          uint `json:",omitempty"`
	ProductName string
	Description string
	Brand       string
}

type ProductItem struct {
	ItemId      uint
	Model       string
	Size        string
	Material    string
	Gender      string
	SKU         string
	QntyInStock int
	Price       int
}
