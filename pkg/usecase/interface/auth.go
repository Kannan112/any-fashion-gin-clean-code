package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
)

type AuthUserCase interface {
	GoogleLoginUser(ctx context.Context, googleuser req.GoogleAuth) (string, string, error)
}
