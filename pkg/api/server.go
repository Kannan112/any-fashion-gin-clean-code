package http

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/kannan112/go-gin-clean-arch/cmd/api/docs"
	handler "github.com/kannan112/go-gin-clean-arch/pkg/api/handler"
	routes "github.com/kannan112/go-gin-clean-arch/pkg/api/routers"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(
	userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler,
	cartHandler *handler.CartHandler,
	productHandler *handler.ProductHandler,
	orderHandler *handler.OrderHandler,
	paymentHandler *handler.PaymentHandler,
	wishlistHandler *handler.WishlistHandler,
	couponHandler *handler.CouponHandler,
	walletHandler *handler.WalletHandler,
	OtpHandler *handler.OtpHandler,
) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	engine.LoadHTMLGlob("./*.html")

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Setup user routes
	routes.SetupUserRoutes(engine, userHandler, cartHandler, productHandler, orderHandler, wishlistHandler, couponHandler, walletHandler, OtpHandler)

	// Setup admin routes
	routes.SetupAdminRoutes(engine, adminHandler, productHandler, orderHandler, couponHandler)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
