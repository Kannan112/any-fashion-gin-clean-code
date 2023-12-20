package repository

import (
	"context"
	"fmt"

	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type RefreshTokenDataBase struct {
	DB *gorm.DB
}

func NewRefreshTokenRepository(DB *gorm.DB) interfaces.RefreshTokenRepository {
	return &RefreshTokenDataBase{DB}
}

func (c *RefreshTokenDataBase) AdminRefreshTokenAdd(token string, adminID uint) error {
	sql := `INSERT INTO admin_refresh_tokens (refresh_token,admin_id)VALUES($1,$2)`
	if err := c.DB.Exec(sql, token, adminID).Error; err != nil {
		return fmt.Errorf("failed to add refresh token")
	}
	return nil
}
func (c *RefreshTokenDataBase) AdminFindRefreshToken(ctx context.Context, adminID uint) (string, error) {
	var token string
	sql := `SELECT refresh_token FROM admin_refresh_tokens where admin_id = $1`
	if err := c.DB.Raw(sql, adminID).Scan(&token).Error; err != nil {
		return "", err
	}
	return token, nil
}

func (c *RefreshTokenDataBase) UserRefreshTokenAdd(token string, userID uint) error {
	var check bool
	tokenExist := `SELECT EXISTS(select * from user_refresh_tokens where users_id=$1)`
	if err := c.DB.Raw(tokenExist, userID).Scan(&check).Error; err != nil {
		return err
	}
	if check {
		updateQuery := `update user_refresh_tokens SET refresh_token=$1 where users_id=$2`
		if err := c.DB.Exec(updateQuery, token, userID).Error; err != nil {
			return err
		}
		return nil
	}
	sql := `INSERT INTO user_refresh_tokens (refresh_token,users_id)VALUES($1,$2)`
	if err := c.DB.Exec(sql, token, userID).Error; err != nil {
		return err
	}
	return nil
}

func (c *RefreshTokenDataBase) UserFindRefreshToken(ctx context.Context, userID uint) (string, error) {
	var token string
	sql := `SELECT refresh_token FROM user_refresh_tokens where users_id = $1`
	if err := c.DB.Raw(sql, userID).Scan(&token).Error; err != nil {
		return "", err
	}
	return token, nil
}
