package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
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
	query := `SELECT EXIST(SELECT 1 FROM users WHERE id=$1)`
	if err := tx.Raw(query, body.UserId).Scan(&exist).Error; err != nil {
		tx.Rollback()

		return err
	}
	if !exist {
		tx.Rollback()
		return fmt.Errorf("no such user")
	}
	if err := tx.Exec("SELECT users SET is_blocked=true WHERE id=?", body.UserId).Error; err != nil {
		tx.Rollback()
		return err
	}
	return fmt.Errorf("irdgfhgh")

}
