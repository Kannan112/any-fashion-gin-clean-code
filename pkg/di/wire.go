//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/kannan112/go-gin-clean-arch/pkg/api"
	"github.com/kannan112/go-gin-clean-arch/pkg/api/handler"
	"github.com/kannan112/go-gin-clean-arch/pkg/config"
	"github.com/kannan112/go-gin-clean-arch/pkg/db"
	"github.com/kannan112/go-gin-clean-arch/pkg/repository"
	"github.com/kannan112/go-gin-clean-arch/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		repository.NewUserRepository,
		repository.NewAdminRepository,
		repository.NewProductRepository,
		repository.NewCartRepository,
		repository.NewOrderRepository,
		repository.NewPaymentRepo,
		repository.NewWishlistRepository,
		repository.NewCouponRepository,
		repository.NewWalletRepository,
		usecase.NewUserUseCase,
		usecase.NewAdminUseCase,
		usecase.NewProductUsecase,
		usecase.NewCartUseCase,
		usecase.NewOrderUseCase,
		usecase.NewPaymentUsecase,
		usecase.NewWishlistUsecase,
		usecase.NewCouponUseCase,
		usecase.NewWalletUseCase,
		handler.NewUserHandler,
		handler.NewAdminSHandler,
		handler.NewProductHandler,
		handler.NewCartHandler,
		handler.NewOrderHandler,
		handler.NewPaymentHandler,
		handler.NewWishlistHandler,
		handler.NewCouponHandler,
		handler.NewWalletHandler,
		//	handler.NewOtpHandler,
		http.NewServerHTTP,
	)

	return &http.ServerHTTP{}, nil
}
