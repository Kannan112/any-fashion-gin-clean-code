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
	GetUserDetailsFromUserID(userId uint) (domain.Users, error)
	AuthSignUp(Oauth req.GoogleAuth) (res.UserResponse, error)
	AuthLogin(email string) (bool, error)
	CheckVerifyPhone(mobileNo string) (bool, error)
	IsSignIn(phno string) (bool, error)
	FindAddress(ctx context.Context, userId int) (bool, error)
	AddAddress(id int, address req.AddAddress) error
	UpdateAddress(id int, addressId int, address req.AddAddress) error
	ListAllAddress(id int) ([]domain.Addresss, error)
	DeleteAddress(ctx context.Context, userId, AddressesId int) ([]domain.Addresss, error)
	ViewProfile(id int) (res.UserData, error)
	EditProfile(id int, profile req.UserReq) (res.UserData, error)
	AccountVerify(phone string) error
}
