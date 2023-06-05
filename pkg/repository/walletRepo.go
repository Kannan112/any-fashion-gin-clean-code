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

func (c *WalletDataBase) ApplyWallet(ctx context.Context, userId uint) error {
	tx := c.DB.Begin()
	var walletCoin float32
	query1 := `select coins from wallets where users_id=$1`
	err := tx.Raw(query1, userId).Scan(&walletCoin).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//collect cart total
	var cartTotal float32
	query2 := `select total from carts where users_id=$1`
	err = tx.Raw(query2, userId).Scan(&cartTotal).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//50rs is fixed min charge
	if walletCoin >= cartTotal {
		walletCoin = cartTotal - 50
	}
	update := `UPDATE carts SET total=total-$1,coin=$2 where users_id=$3`
	err = tx.Exec(update, walletCoin, walletCoin, userId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//update wallet
	query3 := `update wallets set coins=coins-$1 where users_id=$2`
	err = tx.Exec(query3, walletCoin, userId).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil

}
