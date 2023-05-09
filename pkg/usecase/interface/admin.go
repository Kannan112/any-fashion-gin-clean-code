package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

type AdminUsecase interface {
	CreateAdmin(ctx context.Context, admin domain.Admin) error
	AdminLogin(admin req.LoginReq) (string, error)
	BlockUser(body req.BlockData,adminId int)error
	UnblockUser(id int) error
	
}	
