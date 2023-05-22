package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
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

func (c *userUseCase) UserLogin(ctx context.Context, user req.LoginReq) (string, error) {
	userData, err := c.userRepo.UserLogin(ctx, user.Email)
	if err != nil {
		return "", err
	}

	if userData.Email == "" {
		return "", fmt.Errorf("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"id":  userData.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return ss, nil
}

// --------------------------Otp Login_-------------
func (c *userUseCase) OtpLogin(phno string) (string, error) {
	id, err := c.userRepo.OtpLogin(phno)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return ss, nil
}

// -------------------AddAddress-----------
func (c *userUseCase) AddAddress(id int, body req.Address) error {
	err := c.userRepo.AddAddress(id, body)
	return err
}

// -------------------UpdateAddress---------------
func (c *userUseCase) UpdateAddress(id int, addressId int, address req.Address) error {
	err := c.userRepo.UpdateAddress(id, addressId, address)
	return err
}

// --------------------ListAddress------------------
func (c *userUseCase) ListallAddress(id int) ([]domain.Addresss, error) {
	list, err := c.userRepo.ListallAddress(id)
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
