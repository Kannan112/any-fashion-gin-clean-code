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

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{DB}
}
func (c *adminDatabase) FindAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error) {

	if c.DB.Raw("SELECT * FROM admins WHERE email=? OR user_name=?", admin.Email, admin.UserName).Scan(&admin).Error != nil {
		return admin, errors.New("faild to find admin")
	}

	return admin, nil
}

func (c *adminDatabase) CreateAdmin(ctx context.Context, admin domain.Admin) error {

	query := `INSERT INTO admins (user_name,email,password)
								  VALUES($1,$2,$3)`

	if c.DB.Exec(query, admin.UserName, admin.Email, admin.Password).Error != nil {
		return errors.New("faild to save admin")
	}
	return nil
}

func (c *adminDatabase) AdminLogin(email string) (domain.Admin, error) {

	var adminData domain.Admin
	err := c.DB.Raw("SELECT * FROM admins WHERE email=$1", email).Scan(&adminData).Error
	return adminData, err
}

func (c *adminDatabase) BlockUser(body req.BlockData, adminId int) error {
	tx := c.DB.Begin()
	var exist bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE id=$1)`
	if err := tx.Raw(query, body.UserId).Scan(&exist).Error; err != nil {
		fmt.Println("testBlockUse")
		tx.Rollback()
		return err
	}
	if !exist {
		tx.Rollback()
		return fmt.Errorf("no such user")
	}
	if err := tx.Exec("UPDATE users SET is_blocked=true WHERE id=?", body.UserId).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil

}
func (c *adminDatabase) UnblockUser(id int) error {
	tx := c.DB.Begin()
	var IsExist bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE id=$1 AND is_blocked=true)`
	err := c.DB.Raw(query, id).Scan(&IsExist).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if !IsExist {
		tx.Rollback()
		return fmt.Errorf("no such user to unblock")
	}
	err = tx.Exec(`UPDATE users SET is_blocked=false WHERE id=$1`, id).Error
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

// admin dashbord
func (c *adminDatabase) GetDashBord(ctx context.Context) (res.AdminDashboard, error) {
	var admindashbord res.AdminDashboard

	//TotalRevenue
	var TotalRevenue int
	var TotalOrders int
	var TotalProductSold sql.NullInt64
	var TotalUsers int
	query := `SELECT sum(order_total)AS TotalRevenu FROM orders`
	err := c.DB.Raw(query).Scan(&TotalRevenue).Error
	if err != nil {
		return admindashbord, err
	}
	query2 := `SELECT count(*)FROM orders`
	err = c.DB.Raw(query2).Scan(&TotalOrders).Error
	if err != nil {
		return admindashbord, err
	}
	query3 := `SELECT sum(quantity)FROM order_items`
	err = c.DB.Raw(query3).Scan(&TotalProductSold).Error
	if err != nil {
		return admindashbord, err
	}
	query4 := `SELECT count(DISTINCT(id))FROM users;`
	err = c.DB.Raw(query4).Scan(&TotalUsers).Error
	if err != nil {
		return admindashbord, err
	}
	admindashbord.TotalUsers = TotalUsers
	admindashbord.TotalRevenue = TotalRevenue
	admindashbord.TotalOrders = TotalOrders
	admindashbord.TotalProductSold = TotalProductSold
	return admindashbord, err
}
func (c *adminDatabase) ListUsers(ctx context.Context) ([]domain.UsersData, error) {
	tx := c.DB.Begin()
	var user []domain.UsersData
	var check bool
	checking := `SELECT EXISTS(SELECT * FROM users)`
	err := tx.Raw(checking).Scan(&check).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if !check {
		tx.Rollback()
		return nil, fmt.Errorf("no user found")
	}
	query := `select * from users`
	err = c.DB.Raw(query).Scan(&user).Error
	if err != nil {
		return user, err
	}
	return user, err
}
func (c *adminDatabase) FindUserByEmail(ctx context.Context, name string) (domain.UsersData, error) {
	var data domain.UsersData
	query := `SELECT * FROM users WHERE name=$1`
	err := c.DB.Raw(query, name).Scan(&data).Error
	return data, err
}
