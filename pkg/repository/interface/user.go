package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

type UserRepository interface {
	UserSignUp(ctx context.Context, user req.UserReq) (res.UserData, error)
	UserLogin(ctx context.Context, email string) (domain.Users, error)
	IsSignIn(phno string) (bool, error)
	OtpLogin(phone string) (int, error)
	AddAddress(id int, address req.Address) error
	ViewProfile(id int) (res.UserData, error)
	EditProfile(id int,profile req.UserReq)(res.UserData,error)
}
