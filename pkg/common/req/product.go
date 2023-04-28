package req

type Category struct {
	Name string `json:"name" validate:"required"`
}

type Product struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Brand       string `json:"brand" validate:"required"`
	Price int `json:"price" validate:"required"`
	Qty int `json:"qty" validate:"required"`
	CategoryId  string `json:"categoryid" validate:"required"`
}

