package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/kannan112/go-gin-clean-arch/pkg/api/middleware/token"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo         interfaces.UserRepository
	refreshTokenRepo interfaces.RefreshTokenRepository
}

func NewUserUseCase(repo interfaces.UserRepository, refreshtToken interfaces.RefreshTokenRepository) services.UserUseCase {
	return &userUseCase{
		userRepo:         repo,
		refreshTokenRepo: refreshtToken,
	}
}

//--------------------------UserSignUp-----------------------

func (c *userUseCase) UserSignUp(ctx context.Context, user req.UserReq) (res.UserData, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return res.UserData{}, err
	}
	user.Password = string(hash)
	userData, err := c.userRepo.UserSignUp(ctx, user)
	return userData, err
}

//-----------------------UserLogin-----------------------

func (c *userUseCase) UserLogin(ctx context.Context, user req.LoginReq) (res.Token, error) {
	var Token res.Token
	userData, err := c.userRepo.UserLogin(ctx, user.Email)

	if err != nil {
		return Token, err
	}
	if userData.Email == "" {
		return Token, fmt.Errorf("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		return Token, err
	}
	check, err := c.userRepo.CheckVerifyPhone(userData.Mobile)
	if err != nil {
		return Token, err
	}
	if !check {
		return Token, errors.New("account is not verified")
	}

	AccessToken, err := token.GenerateAccessToken(int(userData.ID), "user")
	RefreshToken, err := token.GenerateRefreshToken(int(userData.ID), "user")
	Token.Access_token = AccessToken
	Token.Refresh_token = RefreshToken
	if err := c.refreshTokenRepo.UserRefreshTokenAdd(Token.Refresh_token, userData.ID); err != nil {

		return Token, err
	}
	return Token, nil
}

// -------------------AddAddress-----------
func (c *userUseCase) AddAddress(id int, body req.AddAddress) error {
	err := c.userRepo.AddAddress(id, body)
	return err
}

// -------------------UpdateAddress---------------
func (c *userUseCase) UpdateAddress(id int, addressId int, address req.AddAddress) error {
	err := c.userRepo.UpdateAddress(id, addressId, address)
	return err
}

// --------------------AddressesFind
func (c *userUseCase) FindAddress(ctx context.Context, userId int) (bool, error) {
	checkbool, err := c.userRepo.FindAddress(ctx, userId)
	return checkbool, err
}

// --------------------ListAddress------------------
func (c *userUseCase) ListallAddress(id int) ([]domain.Addresss, error) {
	list, err := c.userRepo.ListAllAddress(id)
	return list, err
}

// -------------------DeletAddresses----------------
func (c *userUseCase) DeleteAddress(ctx context.Context, userId, AddressesId int) ([]domain.Addresss, error) {
	list, err := c.userRepo.DeleteAddress(ctx, userId, AddressesId)
	return list, err
}

// -------------------otp----------------------
func (c *userUseCase) IsSignIn(phone string) (bool, error) {
	signIn, err := c.userRepo.IsSignIn(phone)
	return signIn, err
}

//--------------------ViewProfile------------------

func (c *userUseCase) ViewProfile(id int) (res.UserData, error) {
	userData, err := c.userRepo.ViewProfile(id)
	return userData, err
}

//--------------------Edit Profile----------------

func (c *userUseCase) EditProfile(id int, UpdateProfile req.UserReq) (res.UserData, error) {
	userdata, err := c.userRepo.EditProfile(id, UpdateProfile)
	return userdata, err
}
