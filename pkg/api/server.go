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
	couponHandler *handler.CouponHandler,
	walletHandler *handler.WalletHandler,
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
		address := user.Group("/address", middleware.UserAuth)
		{
			address.POST("add", userHandler.AddAddress)
			address.PATCH("update/:addressId", userHandler.UpdateAddress)
			address.GET("list", userHandler.ListallAddress)
			address.DELETE("delete/:id", userHandler.DeleteAddress)

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
		}
		productitem := user.Group("/product-item")
		{
			productitem.GET("display/:id", productHandler.DisaplyaAllProductItems)
		}
		//cart
		cart := user.Group("/cart", middleware.UserAuth)
		{
			cart.POST("add/:product_items_id", cartHandler.AddToCart)
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
		//coupon
		coupon := user.Group("/coupon", middleware.UserAuth)
		{
			coupon.GET("apply", couponHandler.ApplyCoupon)
		}
		wallet := user.Group("/wallet", middleware.UserAuth)
		{
			wallet.GET("", walletHandler.WallerProfile)
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
			adminUse.GET("all", adminHandler.ListUsers)
			adminUse.GET("email", adminHandler.FindUserByEmail)
			adminUse.PATCH("block", adminHandler.BlockUser)
			adminUse.PATCH("unblock/:userId", adminHandler.UnblockUser)
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
			category.GET("/:category_id", productHandler.DisplayCategory)
		}
		//product
		product := admin.Group("/product", middleware.AdminAuth)
		{
			product.POST("add", productHandler.AddProduct)
			product.PATCH("update/:id", productHandler.UpdateProduct)
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
		paymentMethod := admin.Group("/payment-method", middleware.AdminAuth)
		{
			paymentMethod.POST("add", paymentHandler.SavePaymentMethod)
			paymentMethod.POST("update/:id", paymentHandler.UpdatePaymentMethod)
			paymentMethod.GET("list")
		}

		//coupon
		coupon := admin.Group("/coupon", middleware.AdminAuth)
		{
			coupon.GET("", couponHandler.ViewCoupon)
			coupon.POST("add", couponHandler.AddCoupon)
			coupon.PATCH("update/:couponId", couponHandler.UpdateCoupon)
			coupon.DELETE("delete/:couponId", couponHandler.DeleteCoupon)
		}
		order := admin.Group("/order", middleware.AdminAuth)
		{
			order.GET("", orderHandler.ListAllOrders)
		}
	}

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
