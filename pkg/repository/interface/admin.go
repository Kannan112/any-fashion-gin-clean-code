package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

type AdminRepository interface {
	//SuperAdmin
	//SuperAdminLogin(ctx context.Context)
	//CheckSuperAdmin(ctx context.Context, email string) (bool, error)

	FindAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error)
	CreateAdmin(ctx context.Context, admin domain.Admin) error
	AdminLogin(email string) (domain.Admin, error)

	//DashBord
	GetDashboard(ctx context.Context) (res.AdminDashboard, error)
	BlockUser(body req.BlockData, adminId int) error
	UnblockUser(id int) error
	ListUsers(ctx context.Context) ([]domain.UsersData, error)
	FindUserByEmail(ctx context.Context, name string) (domain.UsersData, error)

	//SalesReport
	ViewSalesReport(ctx context.Context) ([]res.SalesReport, error)
}
