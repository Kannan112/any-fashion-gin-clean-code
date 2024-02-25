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

// @title AnyFashion Application API
// @version 1.0.0
// @description  Backend API built with Golang using Clean Code architecture
// @contact.name API Support
// @contact.email				abhinandarun11@gmail.com
// @securityDefinitions.apikey BearerTokenAuth
// @in header
// @name Authorization
// @BasePath /
// @query.collection.format multi
func NewServerHTTP(
	userHandler *handler.UserHandler, adminHandler *handler.AdminHandler,
	cartHandler *handler.CartHandler, productHandler *handler.ProductHandler,
	orderHandler *handler.OrderHandler, paymentHandler *handler.PaymentHandler,
	wishlistHandler *handler.WishlistHandler, couponHandler *handler.CouponHandler,
	walletHandler *handler.WalletHandler, OtpHandler *handler.OtpHandler,
	RenewHandler *handler.RenewHandler, AuthHandler *handler.AuthHandler,
) *ServerHTTP {

	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	engine.LoadHTMLGlob("./view/*.html")

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// Setup user routes
	routes.SetupUserRoutes(engine.Group("/api"), userHandler, cartHandler, productHandler, orderHandler,
		wishlistHandler, couponHandler, walletHandler, OtpHandler, RenewHandler, AuthHandler, paymentHandler)

	// Setup admin routes
	routes.SetupAdminRoutes(engine.Group("/api"), adminHandler, productHandler, orderHandler, couponHandler, paymentHandler)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8000")
}
