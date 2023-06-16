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

//	type ProductItem struct {
//		Id          uint
//		ItemId      uint
//		Model       string
//		Size        string
//		Material    string
//		Gender      string
//		Sku         string
//		QntyInStock int
//		Price       int
//	}
type OfferTable struct {
	ProductId   uint `gorm:"not null" json:"product_id" validate:"required"`
	Discount    float32
	StartDate   string
	EndDate     string
	Discription string
}

type ProductItem struct {
	ProductId uint
	Sku       string
	Qty       int
	Gender    string
	Model     string
	Size      int
	Color     string
	Material  string
	Price     float64
}
