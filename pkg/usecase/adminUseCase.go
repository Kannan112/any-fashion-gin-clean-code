package usecase

import (
	"context"
	"fmt"

	"github.com/kannan112/go-gin-clean-arch/pkg/api/middleware/token"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepo        interfaces.AdminRepository
	refreshTokenRepo interfaces.RefreshTokenRepository
}

func NewAdminUseCase(adminRepo interfaces.AdminRepository, refresh_token interfaces.RefreshTokenRepository) services.AdminUsecase {
	return &adminUseCase{
		adminRepo:        adminRepo,
		refreshTokenRepo: refresh_token,
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

func (c *adminUseCase) AdminLogin(admin req.LoginReq) (res.Token, error) {
	var result res.Token
	adminData, err := c.adminRepo.AdminLogin(admin.Email)
	if err != nil {
		return result, err
	}

	if adminData.Email == "" {
		return result, fmt.Errorf("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(adminData.Password), []byte(admin.Password))
	if err != nil {
		return result, err
	}

	accessToken, err := token.JWTAccessTokenGen(int(adminData.ID), "admin")
	if err != nil {
		return result, err
	}
	refreshToken, err := token.JWTRefreshTokenGen(int(adminData.ID), "admin")
	if err != nil {

		return result, err
	}
	result = res.Token{
		Access_token:  accessToken,
		Refresh_token: refreshToken,
	}
	if err := c.refreshTokenRepo.AdminRefreshTokenAdd(result.Refresh_token, adminData.ID); err != nil {
		return result, err
	}

	return result, nil
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
