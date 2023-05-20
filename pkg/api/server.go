package http

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/kannan112/go-gin-clean-arch/cmd/api/docs"
	handler "github.com/kannan112/go-gin-clean-arch/pkg/api/handler"
	"github.com/kannan112/go-gin-clean-arch/pkg/api/middleware"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler,

	adminHandler *handler.AdminHandler,
	cartHandler *handler.CartHandler,
	productHandler *handler.ProductHandler,
	orderHandler *handler.OrderHandler,
	paymentHandler *handler.PaymentHandler,
	wishlistHandler *handler.WishlistHandler,
) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	engine.LoadHTMLGlob("./*.html")

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	user := engine.Group("/user")
	{
		user.POST("/signup", userHandler.UserSignUp)
		user.POST("/login", userHandler.UserLogin)
		user.POST("/logout", userHandler.UserLogout)

		//otp
		otp := engine.Group("/otp")
		var OtpHandler handler.OtpHandler
		{
			otp.POST("send", OtpHandler.SendOtp)
			otp.POST("verify", OtpHandler.ValidateOtp)
		}

		//profile
		profile := user.Group("/profile")
		{
			profile.GET("view", middleware.UserAuth, userHandler.ViewProfile)
			profile.PATCH("edit", middleware.UserAuth, userHandler.EditProfile)
		}

		//address
		address := user.Group("/address")
		{
			address.POST("add", middleware.UserAuth, userHandler.AddAddress)
			address.PATCH("update/:addressId", middleware.UserAuth, userHandler.UpdateAddress)
			address.GET("list", middleware.UserAuth, userHandler.ListallAddress)

		}
		//wishlist
		wishlist := user.Group("/wishlist", middleware.UserAuth)
		{
			wishlist.POST("add/:itemId", wishlistHandler.AddToWishlist)
			wishlist.DELETE("remove/:itemId", wishlistHandler.RemoveFromWishlist)
			wishlist.GET("list", wishlistHandler.ListAllWishlist)

		}
		//categories
		categories := user.Group("categories", middleware.UserAuth)
		{
			categories.GET("listall", productHandler.ListCategories)
			categories.GET("listspecific/:category_id", productHandler.DisplayCategory)
		}
		//products
		product := user.Group("product", middleware.UserAuth)
		{
			product.GET("list", productHandler.ListProducts)
			product.GET("list/:id", productHandler.DisplayProduct)
			product.GET("list/:id/:page", productHandler.DisplayProduct)
			product.GET("list/:id/:page/:size", productHandler.DisplayProduct)
			product.GET("list/:id/:page/:size/:sort", productHandler.DisplayProduct)
			product.GET("list/:id/:page/:size/:sort/:category", productHandler.DisplayProduct)
			product.GET("list/:id/:page/:size/:sort/:category/:price", productHandler.DisplayProduct)
		}
		productitem := user.Group("/product-item")
		{
			productitem.GET("display/:id", productHandler.DisaplyaAllProductItems)
		}
		//cart
		cart := user.Group("/cart", middleware.UserAuth)
		{
			cart.POST("add/:product_item_id", cartHandler.AddToCart)
			cart.PATCH("remove/:product_item_id", cartHandler.RemoveFromCart)
			cart.GET("list", cartHandler.ListCart)
		}
		//order
		order := user.Group("/order", middleware.UserAuth)
		{
			order.GET("/razorpay/checkout/:payment_id", orderHandler.RazorPayCheckout)
			order.POST("/razorpay/verify", orderHandler.RazorPayVerify)
			order.POST("orderAll", orderHandler.OrderAll)
			order.PATCH("cancel/:orderId", orderHandler.UserCancelOrder)
			order.GET("listall", orderHandler.ListAllOrders)
		}
	}
	admin := engine.Group("/admin")
	{
		admin.POST("createadmin", adminHandler.CreateAdmin)
		admin.POST("adminlogin", adminHandler.AdminLogin)
		admin.POST("logout", adminHandler.AdminLogout)

		//admin block unblock users
		adminUse := admin.Group("/user", middleware.AdminAuth)
		{
			adminUse.PATCH("block", adminHandler.BlockUser)
			adminUse.PATCH("unblock", adminHandler.UnblockUser)
		}

		//admin dashbord
		DashBord := admin.Group("/dashbord")
		{
			DashBord.GET("list", adminHandler.GetDashBord)
		}

		//categories
		category := admin.Group("/category", middleware.AdminAuth)
		{
			category.POST("add", productHandler.CreateCategory)
			category.PATCH("update/:id", productHandler.UpdatCategory)
			category.DELETE("delete/:category_id")
			category.GET("listall", productHandler.ListCategories)
			category.GET("find/:id", productHandler.DisplayCategory)
		}
		//product
		product := admin.Group("/product", middleware.AdminAuth)
		{
			product.POST("add", productHandler.AddProduct)
			product.PATCH("update", productHandler.UpdateProduct)
		}
		//product item
		productItem := admin.Group("/product-item", middleware.AdminAuth)
		{
			productItem.POST("add", productHandler.AddProductItem)
			productItem.PATCH("update/:id", productHandler.UpdateProductItem)
			productItem.DELETE("delete/:id", productHandler.DeleteProductItem)
			productItem.GET("display/:id", productHandler.DisaplyaAllProductItems)
		}

		//payment-method
		paymentMethod := admin.Group("/payment-method")
		{
			paymentMethod.POST("add", paymentHandler.SavePaymentMethod)
			paymentMethod.POST("update/:id", paymentHandler.UpdatePaymentMethod)
			paymentMethod.GET("list")
		}
	}

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
