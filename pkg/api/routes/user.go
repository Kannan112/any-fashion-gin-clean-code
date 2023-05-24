package routes

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/kannan112/go-gin-clean-arch/pkg/api/handler"
// 	"github.com/kannan112/go-gin-clean-arch/pkg/api/middleware"
// )

// func UserRoutes(
// 	api *gin.RouterGroup,
// 	userHandler *handler.UserHandler,
// 	productHandler *handler.ProductHandler,
// 	cartHandler *handler.CartHandler,
// 	orderHandler *handler.OrderHandler,
// 	otpHandler *handler.OtpHandler,
// 	paymentHandler *handler.PaymentHandler,
// 	wishlistHandler *handler.WishlistHandler,
// ) {

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
