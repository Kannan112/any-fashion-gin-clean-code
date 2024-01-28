package handler

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	handlerUtil "github.com/kannan112/go-gin-clean-arch/pkg/api/handlerUril"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type AdminHandler struct {
	adminUseCase services.AdminUsecase
}

func NewAdminSHandler(admiUseCase services.AdminUsecase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: admiUseCase,
	}
}

// CreateAdmin
// @Summary Create a new admin from admin panel
// @ID CreateAdmin
// @Description admin creation
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin body req.CreateAdmin true "New Admin details"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /api/admin/createadmin [post]
func (cr *AdminHandler) CreateAdmin(c *gin.Context) {
	var adminData req.CreateAdmin
	if err := c.Bind(&adminData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, res.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	adminDetails, err := cr.adminUseCase.CreateAdmin(c.Request.Context(), adminData)

	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Can't Create Admin",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, res.Response{
		StatusCode: 200,
		Message:    "Admin created",
		Data:       adminDetails,
		Errors:     nil,
	})

}

// @Summary Admin Login
// @Description Logs in an admin user and returns an authentication token
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin body req.LoginReq true "Admin login details"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /api/admin/adminlogin [post]
func (cr *AdminHandler) AdminLogin(c *gin.Context) {
	var admin req.LoginReq
	err := c.Bind(&admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	result, err := cr.adminUseCase.AdminLogin(admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "logined success fully",
		Data: res.AdminToken{
			Token: result.Access_token,
			//Refresh_token: result.Refresh_token,
		},
		Errors: nil,
	})
}

// Admin Logout
// @Summary Admin Logout
// @ID admin-logout
// @Description Logs out a logged-in admin from the E-commerce web api admin panel
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400
// @Router /api/admin/logout [post]
func (cr *AdminHandler) AdminLogout(c *gin.Context) {
	c.SetCookie("AdminAuth", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "admin logouted",
		Data:       nil,
		Errors:     nil,
	})

}

// @Summary Admin BlockUser
// @Description admin block user access to the store
// @Tags Admin
// @Accept json
// @Produce json
// @Param blocking_details body req.BlockData true "User bolocking details"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Security BearerTokenAuth
// @Router /api/admin/user/block [patch]
func (cr *AdminHandler) BlockUser(c *gin.Context) {
	var body req.BlockData
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "binding fail",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	adminId, errf := handlerUtil.GetAdminIdFromContext(c)
	if errf != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to find admin_id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.adminUseCase.BlockUser(body, adminId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Can't Block",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "Blocked user",
		Data:       nil,
		Errors:     nil,
	})

}

// UnBlockUser
// @Summary Admin can unbolock a blocked user
// @ID unblock-users
// @Description Admins can block users
// @Tags Admin
// @Accept json
// @Produce json
// @Param user_id path string true "ID of the user to be blocked"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Security BearerTokenAuth
// @Router /api/admin/user/unblock/{user_id} [patch]
func (cr *AdminHandler) UnblockUser(c *gin.Context) {
	paramsId := c.Param("userId")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "bind failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.adminUseCase.UnblockUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant unblock user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "user unblocked",
		Data:       nil,
		Errors:     nil,
	})
}

// AdminDashboard
// @Summary Admin Dashboard
// @ID admin-dashboard
// @Description Admin can access dashboard and view details regarding orders, products, etc.
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Security BearerTokenAuth
// @Router /api/admin/dashbord/list [get]
func (cr *AdminHandler) GetDashBord(c *gin.Context) {
	data, err := cr.adminUseCase.GetDashBord(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to show",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "success",
		Data:       data,
		Errors:     nil,
	})
}

func (c *AdminHandler) ListUsers(ctx *gin.Context) {
	data, err := c.adminUseCase.ListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to list",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "List of users",
		Data:       data,
		Errors:     nil,
	})
}

func (c *AdminHandler) FindUserByEmail(ctx *gin.Context) {
	var userEmail req.UserEmail
	err := ctx.BindJSON(&userEmail)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Failed to bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	fmt.Println("EmailFind", userEmail.Email)
	data, err := c.adminUseCase.FindUserByEmail(ctx, userEmail.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Failed find user",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "user found",
		Data:       data,
		Errors:     nil,
	})

}

func (cr *AdminHandler) ViewSalesReport(ctx *gin.Context) {
	sales, err := cr.adminUseCase.ViewSalesReport(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant get sales report",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "Sales report",
		Data:       sales,
		Errors:     nil,
	})

}

// DownloadSalesReport
// @Summary Admin can download sales report
// @ID download-sales-report
// @Description Admin can download sales report in .csv format
// @Tags Admin
// @Accept json
// @Produce json
// @Failure 400 {object} res.Response
// @Security BearerTokenAuth
// @Router /api/admin/sales/download [get]
func (cr *AdminHandler) DownloadSalesReport(ctx *gin.Context) {
	sales, err := cr.adminUseCase.ViewSalesReport(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant get sales report",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	// Set headers so browser will download the file
	ctx.Header("Content-Type", "text/csv")
	ctx.Header("Content-Disposition", "attachment;filename=sales.csv")

	// Create a CSV writer using our response writer as our io.Writer
	wr := csv.NewWriter(ctx.Writer)

	// Write CSV header row
	headers := []string{"Name", "Mobile", "OrderStatus", "OrderTime", "OrderTotal"}
	if err := wr.Write(headers); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Write data rows
	for _, sale := range sales {
		row := []string{sale.Name, sale.Mobile, sale.OrderStatus, sale.OrderTime.Format("2006-01-02 15:04:05"), strconv.Itoa(sale.OrderTotal)}
		if err := wr.Write(row); err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	// Flush the writer's buffer to ensure all data is written to the client
	wr.Flush()
	if err := wr.Error(); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

}
