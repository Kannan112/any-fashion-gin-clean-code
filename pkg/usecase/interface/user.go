package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
)

type UserUseCase interface {
	UserSignUp(ctx context.Context, user req.UserReq) (res.UserData, error)
	UserLogin(ctx context.Context, user req.LoginReq) (string, error)
	IsSignIn(phno string) (bool, error)
	OtpLogin(phone string) (string, error)
	AddAddress(id int, body req.Address) error
	ViewProfile(id int) (res.UserData, error)
	EditProfile(id int,UpdateProfile req.UserReq)(res.UserData,error)
}
