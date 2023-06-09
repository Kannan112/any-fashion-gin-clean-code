package res

// /response needed only
type OrderDetails struct {
	AddressId   uint
	HouseNumber int
	Street      string
	City        string
	District    string
	Landmark    string
	Pincode     int
	OrderTime   string
	OrderStatus string
}

type UserOrder struct {
	Name          string
	Mobile        string
	ProductItemId uint
	AddressId     uint
	HouseNumber   int
	Street        string
	City          string
	District      string
	Landmark      string
	Pincode       int
	OrderTime     string
	OrderStatus   string
}

type OrderData struct {
	Name        string
	Mobile      string
	Price       int
	ProductName string
	Brand       string
	HouseNumber int
	Street      string
	City        string
	District    string
	Landmark    string
	Pincode     int
}
