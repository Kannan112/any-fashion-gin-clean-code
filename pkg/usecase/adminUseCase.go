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

type adminUseCase struct {
	adminRepo interfaces.AdminRepository
}

func NewAdminUseCase(adminRepo interfaces.AdminRepository) services.AdminUsecase {
	return &adminUseCase{
		adminRepo: adminRepo,
	}
}

func (c *adminUseCase) CreateAdmin(ctx context.Context, admin req.CreateAdmin) (res.AdminData, error) {
	// IsSuper, err := c.adminRepo.IsSuperAdmin(createrId)
	// if err != nil {
	// 	return res.AdminData{}, err
	// }
	// if !IsSuper {
	// 	return res.AdminData{}, fmt.Errorf("not a super admin")
	// }

	hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return res.AdminData{}, err
	}
	admin.Password = string(hash)
	adminData, err := c.adminRepo.CreateAdmin(admin)

	return adminData, err
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

//admin dashbord

func (c *adminUseCase) GetDashBord(ctx context.Context) (res.AdminDashboard, error) {
	data, err := c.adminRepo.GetDashboard(ctx)
	return data, err

}
func (c *adminUseCase) ListUsers(ctx context.Context) ([]domain.UsersData, error) {
	data, err := c.adminRepo.ListUsers(ctx)
	return data, err
}
func (c *adminUseCase) FindUserByEmail(ctx context.Context, name string) (domain.UsersData, error) {
	data, err := c.adminRepo.FindUserByEmail(ctx, name)
	return data, err
}
func (c *adminUseCase) ViewSalesReport(ctx context.Context) ([]res.SalesReport, error) {
	report, err := c.adminRepo.ViewSalesReport(ctx)
	return report, err
}
