package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type AdminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &AdminDatabase{DB}
}

// func (c *AdminDatabase) CheckSuperAdmin(ctx context.Context, email string) (bool, error) {
// 	check := `select exists(select * from admins where email AND is_super=true)`
// 	err := c.DB.Exec(check, email).Error
// 	return true, err
// }

func (c *AdminDatabase) IsSuperAdmin(createrId int) (bool, error) {
	var isSuper bool
	query := "SELECT is_super_admin FROM admins WHERE id=$1"
	err := c.DB.Raw(query, createrId).Scan(&isSuper).Error
	return isSuper, err
}

// Find Admin FOR super admin
func (a *AdminDatabase) FindAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error) {
	if err := a.DB.Where("email = ? OR user_name = ?", admin.Email, admin.UserName).First(&admin).Error; err != nil {
		return admin, errors.New("failed to find admin")
	}

	return admin, nil
}

// Create admin
func (a *AdminDatabase) CreateAdmin(admin req.CreateAdmin) (res.AdminData, error) {
	var adminData res.AdminData
	// if err := a.DB.Create(&admin).Error; err != nil {
	// 	return adminData, errors.New("failed to save admin")
	// }
	addQuery := `INSERT INTO admins (user_name,email,password) VALUES($1,$2,$3)`
	err := a.DB.Raw(addQuery, admin.Name, admin.Email, admin.Password).Scan(&adminData).Error
	return adminData, err
}

// Admin login
func (a *AdminDatabase) AdminLogin(email string) (domain.Admin, error) {
	var adminData domain.Admin
	err := a.DB.Where("email = ?", email).First(&adminData).Error
	return adminData, err
}

// Admin block user
func (a *AdminDatabase) BlockUser(body req.BlockData, adminID int) error {
	tx := a.DB.Begin()

	var exist bool
	if err := tx.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)", body.UserId).Scan(&exist).Error; err != nil {
		tx.Rollback()
		return err
	}

	if !exist {
		tx.Rollback()
		return fmt.Errorf("no such user")
	}

	if err := tx.Exec("UPDATE users SET is_blocked = true WHERE id = ?", body.UserId).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Unblock user by admin
func (a *AdminDatabase) UnblockUser(id int) error {
	tx := a.DB.Begin()

	var isExist bool
	err := tx.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = ? AND is_blocked = true)", id).Scan(&isExist).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if !isExist {
		tx.Rollback()
		return fmt.Errorf("no such user to unblock")
	}

	err = tx.Exec("UPDATE users SET is_blocked = false WHERE id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Admin dashboard { Showing : TOTAL SALES,TOTAL REVENUE,TOTAL ORDERS,TOTAL PRODUCT SOLD }
func (a *AdminDatabase) GetDashboard(ctx context.Context) (res.AdminDashboard, error) {
	var adminDashboard res.AdminDashboard
	var totalRevenue int
	var totalOrders int
	var totalProductSold sql.NullInt64
	var totalUsers int

	if err := a.DB.Raw("SELECT SUM(order_total) AS TotalRevenue FROM orders").Scan(&totalRevenue).Error; err != nil {
		return adminDashboard, err
	}

	if err := a.DB.Raw("SELECT COUNT(*) FROM orders").Scan(&totalOrders).Error; err != nil {
		return adminDashboard, err
	}

	if err := a.DB.Raw("SELECT SUM(quantity) FROM order_items").Scan(&totalProductSold).Error; err != nil {
		return adminDashboard, err
	}

	if err := a.DB.Raw("SELECT COUNT(DISTINCT id) FROM users").Scan(&totalUsers).Error; err != nil {
		return adminDashboard, err
	}

	adminDashboard.TotalUsers = totalUsers
	adminDashboard.TotalRevenue = totalRevenue
	adminDashboard.TotalOrders = totalOrders
	adminDashboard.TotalProductSold = totalProductSold

	return adminDashboard, nil
}

// List all users by admin
func (a *AdminDatabase) ListUsers(ctx context.Context) ([]domain.UsersData, error) {
	tx := a.DB.Begin()

	var users []domain.UsersData
	var check bool

	if err := tx.Raw("SELECT EXISTS(SELECT * FROM users)").Scan(&check).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if !check {
		tx.Rollback()
		return nil, fmt.Errorf("no user found")
	}

	if err := a.DB.Raw("SELECT * FROM users").Scan(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

// Search a specific user by email or name or phone number
func (a *AdminDatabase) FindUserByEmail(ctx context.Context, email string) (domain.UsersData, error) {
	var data domain.UsersData
	if err := a.DB.Where("email = ?", email).First(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (c *AdminDatabase) ViewSalesReport(ctx context.Context) ([]res.SalesReport, error) {
	var sales []res.SalesReport
	if err := c.DB.Raw("select * from orders o join users u on u.id=o.users_id").Scan(&sales).Error; err != nil {
		return sales, err
	}
	return sales, nil
}
