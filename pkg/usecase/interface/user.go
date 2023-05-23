package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

type UserUseCase interface {
	UserSignUp(ctx context.Context, user req.UserReq) (res.UserData, error)
	UserLogin(ctx context.Context, user req.LoginReq) (string, error)
	IsSignIn(phno string) (bool, error)
	OtpLogin(phone string) (string, error)
	AddAddress(id int, body req.Address) error
	UpdateAddress(id int, addressId int, address req.Address) error
	ListallAddress(id int) ([]domain.Addresss, error)
	DeleteAddress(ctx context.Context, userId, AddressesId int) ([]domain.Addresss, error)
	ViewProfile(id int) (res.UserData, error)
	EditProfile(id int, UpdateProfile req.UserReq) (res.UserData, error)

}
