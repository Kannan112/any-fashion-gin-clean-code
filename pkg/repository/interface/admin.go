package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

type AdminRepository interface {
	FindAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error)
	CreateAdmin(ctx context.Context, admin domain.Admin) error
	AdminLogin(email string) (domain.Admin, error)
	BlockUser(body req.BlockData, adminId int) error
	UnblockUser(id int) error

}
