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

		// Cart Items
		cartitem := user.Group("/cart-item", middleware.UserAuth)
		{
			cartitem.GET("list", cartHandler.ListCartItems)
			// 	cartitem.GET("list/:id", cartHandler.DisplayCartItem)
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
			coupon.PATCH("remove", couponHandler.RemoveCoupon)
		}

		// Wallet
		wallet := user.Group("/wallet", middleware.UserAuth)
		{
			wallet.GET("", walletHandler.WallerProfile)
			//wallet apply while purchasing{reduce the amount in wallet}
		}
	}
}
