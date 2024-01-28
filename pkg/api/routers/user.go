package http

import (
	"github.com/gin-gonic/gin"
	handler "github.com/kannan112/go-gin-clean-arch/pkg/api/handler"
	"github.com/kannan112/go-gin-clean-arch/pkg/api/middleware"
)

func SetupUserRoutes(engine *gin.RouterGroup, userHandler *handler.UserHandler, cartHandler *handler.CartHandler, productHandler *handler.ProductHandler, orderHandler *handler.OrderHandler, wishlistHandler *handler.WishlistHandler, couponHandler *handler.CouponHandler, walletHandler *handler.WalletHandler, otpHandler *handler.OtpHandler, renew *handler.RenewHandler, authHandler *handler.AuthHandler, paymentHandler *handler.PaymentHandler) {

	engine.POST("/renew-token", renew.GetAccessToken)
	auth := engine.Group("/auth")
	{
		auth.GET("/google-login", authHandler.GoogleLogin)
		auth.GET("/google-callback", authHandler.GoogleAuthCallback)
	}
	user := engine.Group("/user")

	user.GET("/login", userHandler.LLLogin)
	{
		//otp
		otp := user.Group("/otp")
		{
			otp.POST("send", otpHandler.SendOtp)
			otp.POST("verify", otpHandler.ValidateOtp)
		}
		// User routes
		user.POST("/signup", userHandler.UserSignUp)
		user.POST("/login", userHandler.UserLogin)
		user.GET("/logout", userHandler.UserLogout)

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
			address.DELETE("delete/:addressId", userHandler.DeleteAddress)
		}

		// Wishlist
		wishlist := user.Group("/wishlist", middleware.UserAuth)
		{
			wishlist.POST("add/:itemId", wishlistHandler.AddToWishlist)
			wishlist.DELETE("remove/:itemId", wishlistHandler.RemoveFromWishlist)
			wishlist.GET("list", wishlistHandler.ListAllWishlist)
		}

		// Categories
		categories := user.Group("category", middleware.UserAuth)
		{
			categories.GET("listall", productHandler.UserListCategory())
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
			productitem.GET("/:product_id", productHandler.DisaplyaAllProductItems)
		}

		cart := user.Group("/cart", middleware.UserAuth)
		{
			cart.POST("add/:product_items_id", cartHandler.AddToCart)
			cart.DELETE("remove/:product_item_id", cartHandler.RemoveFromCart)
			cart.GET("list", cartHandler.ListCart)

		}

		// Cart Items
		cartitem := user.Group("/cart-item", middleware.UserAuth)
		{
			cartitem.GET("list", cartHandler.ListCartItems)
			//	cartitem.GET("list/:id", cartHandler.DisplayCartItem)
		}

		// Order
		order := user.Group("/order", middleware.UserAuth)
		{

			order.GET("orderall", orderHandler.OrderAll)
			order.PATCH("cancel/:orderId", orderHandler.UserCancelOrder)
			order.GET("listall", orderHandler.ListOrdersOfUsers)
			order.GET("/:orderId", orderHandler.OrderDetails)
		}
		//	engine.GET("/razorpay/checkout/:payment_id", paymentHandler.RazorPayCheckout)
		//	engine.POST("/razorpay/verify", paymentHa ndler.RazorPayVerify)

		// Coupon
		coupon := user.Group("/coupon", middleware.UserAuth)
		{
			coupon.POST("apply", couponHandler.ApplyCoupon)
			coupon.PATCH("remove", couponHandler.RemoveCoupon)
		}

		// Wallet
		wallet := user.Group("/wallet", middleware.UserAuth)
		{
			wallet.GET("", walletHandler.WallerProfile)
			wallet.POST("/apply", walletHandler.ApplyWallet)
			wallet.DELETE("/remove", walletHandler.RemoveWallet)
			//wallet apply while purchasing{reduce the amount in wallet}
		}

		payment := user.Group("/payment", middleware.UserAuthCookie)
		{
			payment.GET("payment-methods", paymentHandler.GetPaymentMethodUser())
			payment.GET("/checkout/payment-select-page", paymentHandler.CartOrderPaymentSelectPage)

			// razorpay-payment
			payment.GET("/razorpay-checkout", paymentHandler.RazorPayCheckout)
			payment.POST("/razorpay-verify", paymentHandler.RazorPayVerify)

			// stripe-payment
			payment.POST("/stripe-checkout", paymentHandler.StripeCheckout)
			payment.POST("/stripe-verify", paymentHandler.StripePaymentVerify)
		}

	}
}
