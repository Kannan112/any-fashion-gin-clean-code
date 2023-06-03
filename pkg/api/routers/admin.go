package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/kannan112/go-gin-clean-arch/pkg/api/handler"
	"github.com/kannan112/go-gin-clean-arch/pkg/api/middleware"
)

func SetupAdminRoutes(engine *gin.Engine, adminHandler *handler.AdminHandler, productHandler *handler.ProductHandler, orderHandler *handler.OrderHandler, couponHandler *handler.CouponHandler) {
	admin := engine.Group("/admin")
	{
		// Admin routes
		admin.POST("createadmin", adminHandler.CreateAdmin)
		admin.POST("adminlogin", adminHandler.AdminLogin)
		admin.POST("logout", adminHandler.AdminLogout)

		// Admin block unblock users
		adminUse := admin.Group("/user", middleware.AdminAuth)
		{
			adminUse.GET("all", adminHandler.ListUsers)
			adminUse.GET("email", adminHandler.FindUserByEmail)
			adminUse.PATCH("block", adminHandler.BlockUser)
			adminUse.PATCH("unblock/:userId", adminHandler.UnblockUser)
		}

		// Admin dashboard
		DashBord := admin.Group("/dashbord")
		{
			DashBord.GET("list", adminHandler.GetDashBord)
		}

		// Categories
		category := admin.Group("/category", middleware.AdminAuth)
		{
			category.POST("add", productHandler.CreateCategory)
			category.PATCH("update/:id", productHandler.UpdatCategory)
			category.DELETE("delete/:category_id")
			category.GET("listall", productHandler.ListCategories)
			category.GET("/:category_id", productHandler.DisplayCategory)
		}

		// Product
		product := admin.Group("/product", middleware.AdminAuth)
		{
			product.POST("add", productHandler.AddProduct)
			product.PATCH("update/:id", productHandler.UpdateProduct)
		}

		// Product item
		productItem := admin.Group("/product-item", middleware.AdminAuth)
		{
			productItem.POST("add", productHandler.AddProductItem)
			productItem.PATCH("update/:id", productHandler.UpdateProductItem)
			productItem.DELETE("delete/:id", productHandler.DeleteProductItem)
			productItem.GET("display/:id", productHandler.DisaplyaAllProductItems)
		}

		// Coupon
		coupon := admin.Group("/coupon", middleware.AdminAuth)
		{
			coupon.GET("", couponHandler.ViewCoupon)
			coupon.POST("add", couponHandler.AddCoupon)
			coupon.PATCH("update/:couponId", couponHandler.UpdateCoupon)
			coupon.DELETE("delete/:couponId", couponHandler.DeleteCoupon)
		}

		// Order
		order := admin.Group("/order", middleware.AdminAuth)
		{
			order.GET("", orderHandler.ViewOrder)
			order.GET("/:orderid", orderHandler.ListAllOrders)
			order.GET("/placed", orderHandler.ListOrderByPlaced)
			order.GET("/cancelled", orderHandler.ListOrderByCancelled)
		}

		// offer side
		offer := admin.Group("offer", middleware.AdminAuth)
		{
			offer.POST("/", productHandler.SaveOffer)
		}
	}
}
