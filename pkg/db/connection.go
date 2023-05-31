package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "github.com/kannan112/go-gin-clean-arch/pkg/config"
	domain "github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(
		&domain.Users{},
		&domain.Admin{},
		&domain.Category{},
		&domain.Product{},
		&domain.ProductItem{},
		&domain.UserInfo{},
		&domain.Address{},
		&domain.Carts{},
		&domain.CartItem{},
		&domain.Orders{},
		&domain.OrderItem{},
		//&domain.OrderStatus{},
		&domain.PaymentMethod{},
		&domain.WishList{},
		&domain.Coupon{},
		&domain.Wallet{},
	)

	return db, dbErr
}
