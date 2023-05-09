package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepo interfaces.AdminRepository
}

func NewAdminUseCase(adminRepo interfaces.AdminRepository) services.AdminUsecase {
	return &adminUseCase{
		adminRepo: adminRepo,
	}
}

func (c *adminUseCase) CreateAdmin(ctx context.Context, admin domain.Admin) error {

	cadmin, err := c.adminRepo.FindAdmin(ctx, admin)
	if err != nil {
		return err
	} else if cadmin.ID != 0 {
		return errors.New("can't save admin already exist")
	}
	// generate a hashed password for admin
	hashPass, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)

	if err != nil {
		return errors.New("failed to generate hashed password for admin")
	}
	// set the hashed password on the admin
	admin.Password = string(hashPass)

	return c.adminRepo.CreateAdmin(ctx, admin)
}

func (c *adminUseCase) AdminLogin(admin req.LoginReq) (string, error) {
	adminData, err := c.adminRepo.AdminLogin(admin.Email)
	if err != nil {
		fmt.Println("second")
		return "", err
	}

	if adminData.Email == "" {
		return "", fmt.Errorf("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(adminData.Password), []byte(admin.Password))
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"id":  adminData.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return ss, nil
}
func (c adminUseCase) BlockUser(body req.BlockData, adminId int) error {
	err := c.adminRepo.BlockUser(body, adminId)
	return err
}
func (c *adminUseCase) UnblockUser(id int) error {
	err := c.adminRepo.UnblockUser(id)
	return err
}
