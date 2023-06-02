package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/kannan112/go-gin-clean-arch/pkg/api/handler"
	"github.com/kannan112/go-gin-clean-arch/pkg/api/middleware"
)

func SetupUserRoutes(engine *gin.Engine, userHandler *handler.UserHandler, cartHandler *handler.CartHandler, productHandler *handler.ProductHandler, orderHandler *handler.OrderHandler, wishlistHandler *handler.WishlistHandler, couponHandler *handler.CouponHandler, walletHandler *handler.WalletHandler, otpHandler *handler.OtpHandler) {
	user := engine.Group("/user")
	{
		// User routes
		user.POST("/signup", userHandler.UserSignUp)
		user.POST("/login", userHandler.UserLogin)
		user.POST("/logout", userHandler.UserLogout)

		// Profile
		profile := user.Group("/profile")
		{
			profile.GET("view", middleware.UserAuth, userHandler.ViewProfile)
			profile.PATCH("edit", middleware.UserAuth, userHandler.EditProfile)
		}

		// Address
		address := user.Group("/address", middleware.UserAuth)
		{
			address.POST("add", userHandler.AddAddress)
			address.PATCH("update/:addressId", userHandler.UpdateAddress)
			address.GET("list", userHandler.ListallAddress)
			address.DELETE("delete/:id", userHandler.DeleteAddress)
		}

		// Wishlist
		wishlist := user.Group("/wishlist", middleware.UserAuth)
		{
			wishlist.POST("add/:itemId", wishlistHandler.AddToWishlist)
			wishlist.DELETE("remove/:itemId", wishlistHandler.RemoveFromWishlist)
			wishlist.GET("list", wishlistHandler.ListAllWishlist)
		}

		// Categories
		categories := user.Group("categories", middleware.UserAuth)
		{
			categories.GET("listall", productHandler.ListCategories)
			categories.GET("listspecific/:category_id", productHandler.DisplayCategory)
		}

		// Products
		product := user.Group("product", middleware.UserAuth)
		{
			product.GET("list", productHandler.ListProducts)
			product.GET("list/:id", productHandler.DisplayProduct)
		}

		productitem := user.Group("/product-item")
		{
			productitem.GET("display/:id", productHandler.DisaplyaAllProductItems)
		}

		// Cart
		cart := user.Group("/cart", middleware.UserAuth)
		{
			cart.POST("add/:product_items_id", cartHandler.AddToCart)
			cart.PATCH("remove/:product_item_id", cartHandler.RemoveFromCart)
			cart.GET("list", cartHandler.ListCart)
		}

		// Order
		order := user.Group("/order", middleware.UserAuth)
		{
			order.GET("/razorpay/checkout/:payment_id", orderHandler.RazorPayCheckout)
			order.POST("/razorpay/verify", orderHandler.RazorPayVerify)
			order.POST("orderAll", orderHandler.OrderAll)
			order.PATCH("cancel/:orderId", orderHandler.UserCancelOrder)
			order.GET("listall", orderHandler.ListAllOrders)
			order.GET("/:orderId", orderHandler.OrderDetails)
		}

		// Coupon
		coupon := user.Group("/coupon", middleware.UserAuth)
		{
			coupon.GET("apply", couponHandler.ApplyCoupon)
		}

		// Wallet
		wallet := user.Group("/wallet", middleware.UserAuth)
		{
			wallet.GET("", walletHandler.WallerProfile)
		}
	}
}

// 	// User routes that don't require authentication
// 	api.POST("/signup", userHandler.UserSignUp)
// 	// api.POST("/login/email", userHandler.LoginWithEmail)
// 	// api.POST("/login/phone", userHandler.)
// 	// api.POST("/send-otp", otpHandler.SendOtp)
// 	api.POST("/verify-otp", otpHandler.ValidateOtp)

// 	// Category routes
// 	category := api.Group("/categories")
// 	{
// 		category.GET("", productHandler.ListCategories)
// 		category.GET("/:id", productHandler.DisplayCategory)
// 	}

// 	// Brand routes
// 	// brand := api.Group("/brands")
// 	// {
// 	// 	brand.GET("", productHandler)
// 	// 	brand.GET("/:id", productHandler.ViewBrandByID)
// 	// }

// 	// Product routes
// 	product := api.Group("/products")
// 	{
// 		product.GET("", productHandler.ListProducts)
// 		product.GET("/:id", productHandler.DisplayProduct)
// 	}

// 	// Product item routes
// 	productItem := api.Group("/product-items")
// 	{
// 		productItem.GET("", productHandler.DisaplyaAllProductItems)
// 		//productItem.GET("/:id", productHandler.)
// 	}

// 	// User routes that require authentication
// 	api.Use(middleware.UserAuth)
// 	{
// 		api.GET("/profile", userHandler.ViewProfile)
// 		api.PATCH("/profile/edit", userHandler.EditProfile)
// 		api.GET("/logout", userHandler.UserLogout)

// 		// Address routes
// 		address := api.Group("/addresses")
// 		{
// 			address.POST("/", userHandler.AddAddress)
// 			address.PUT("/", userHandler.UpdateAddress)
// 			address.DELETE("/:id",userHandler.DeleteAddress)
// 			address.GET("")
// 		}

// 		// Cart routes
// 		cart := api.Group("/cart")
// 		{
// 			cart.POST("/add/:product_item_id", cartHandler.AddToCart)
// 			cart.DELETE("/remove/:product_item_id", cartHandler.RemoveFromCart)
// 			cart.POST("/coupon/:coupon_id", cartHandler.AddCouponToCart)
// 			cart.GET("", cartHandler.ViewCart)
// 			cart.DELETE("", cartHandler.EmptyCart)
// 		}

// 		// Coupon routes
// 		coupon := api.Group("/coupons")
// 		{
// 			coupon.GET("", productHandler.ViewAllCoupons)
// 			coupon.GET("/:id", productHandler.ViewCouponByID)
// 		}

// 		// Order routes
// 		order := api.Group("/orders")
// 		{
// 			order.POST("", orderHandler.BuyProductItem)
// 			order.POST("/buy-all", orderHandler.BuyAll)
// 			order.GET("/:id", orderHandler.ViewOrderByID)
// 			order.GET("", orderHandler.ViewAllOrders)
// 			order.PUT("/cancel/:id", orderHandler.CancelOrder)
// 			order.POST("/return", orderHandler.ReturnRequest)
// 		}

// 		// Payment routes
// 		payment := api.Group("/payments")
// 		{
// 			payment.GET("/razorpay/:order_id", paymentHandler.CreateRazorpayPayment)
// 			payment.GET("/success", paymentHandler.PaymentSuccess)
// 		}

// 		//wishlist routes
// 		wishlist := api.Group("/wishlist")
// 		{
// 			wishlist.GET("/", wishlistHandler.ViewWishlist)
// 			wishlist.POST("/:id", wishlistHandler.AddToWishlist)
// 			wishlist.DELETE("/:id", wishlistHandler.RemoveFromWishlist)
// 			wishlist.DELETE("/", wishlistHandler.EmptyWishlist)
// 		}
// 	}

// }
