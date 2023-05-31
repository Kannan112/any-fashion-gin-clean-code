package repository

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	"gorm.io/gorm"
)

type WalletDataBase struct {
	DB *gorm.DB
}

func NewWalletRepository(DB *gorm.DB) interfaces.WalletRepo {
	return &WalletDataBase{
		DB: DB,
	}
}
func (c *WalletDataBase) CollentTheRefundAmount(ctx context.Context, OrderId int) (float32, error) {
	var price float32
	Amount := `SELECT price FROM orders WHERE order_id=$1`
	err := c.DB.Raw(Amount, OrderId).Scan(&price).Error
	return price, err
}

func (c *WalletDataBase) SaveWallet(ctx context.Context, UserID int) error {
	tx := c.DB.Begin()
	query := `INSERT INTO wallets (users_id,coins) VALUES($1,0)`
	err := c.DB.Exec(query, UserID).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return err
}
func (c *WalletDataBase) AddCoinToWallet(ctx context.Context, price float32, usersId uint) error {
	addToWallet := `UPDATE wallets SET coins=coins+$1 WHERE users_id=$2`
	err := c.DB.Exec(addToWallet, price, usersId).Error
	return err
}

func (c *WalletDataBase) WallerProfile(ctx context.Context, userID uint) (res.Wallet, error) {
	var profile res.Wallet
	walletProfile := `SELECT users_id,coins FROM wallets WHERE users_id=$1`
	err := c.DB.Raw(walletProfile, userID).Scan(&profile).Error
	return profile, err
}
