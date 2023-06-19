package res

import (
	"database/sql"
	"time"
)

type AdminData struct {
	Id    int
	Name  string
	Email string
}
type AdminDashboard struct {
	TotalUsers       int
	TotalOrders      int
	TotalProductSold sql.NullInt64
	TotalRevenue     sql.NullFloat64
}
type SalesReport struct {
	Id          int
	Name        string
	Mobile      string
	OrderStatus string
	OrderTime   time.Time
	OrderTotal  int
}
