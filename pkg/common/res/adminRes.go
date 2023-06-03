package res

import "database/sql"

type AdminData struct {
	Id    int
	Name  string
	Email string
}
type AdminDashboard struct {
	TotalUsers       int
	TotalOrders      int
	TotalProductSold sql.NullInt64
	TotalRevenue     int
}
