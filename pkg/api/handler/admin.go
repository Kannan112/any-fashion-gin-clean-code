package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
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
// @Param admin_details body req.CreateAdmin true "New Admin details"
// @Success 201 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /admin/createadmin [post]

func (cr *AdminHandler) CreateAdmin(c *gin.Context) {
	var adminData domain.Admin
	if err := c.Bind(&adminData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, res.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	err := cr.adminUseCase.CreateAdmin(c.Request.Context(), adminData)

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
		StatusCode: 201,
		Message:    "Admin created",
		Data:       nil,
		Errors:     nil,
	})

}

// AdminLogin
// @Summary Admin Login
// @ID admin-login
// @Description Admin login
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin_credentials body req.LoginReq true "Admin login credentials"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /admin/adminlogin [post]
func (cr *AdminHandler) AdminLogin(c *gin.Context) {
	var admin req.LoginReq
	fmt.Println("1")
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
	ss, err := cr.adminUseCase.AdminLogin(admin)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("AdminAuth", ss, 3660*24*30, "", "", false, true)
	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "logined success fully",
		Data:       nil,
		Errors:     nil,
	})
}

// AdminLogout
// @Summary Admin Logout
// @ID adminlogout
// @Description Logs out a logged-in admin from the E-commerce web api admin panel
// @Tags Admin
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400
// @Router /admin/adminlogout [post]

func (cr *AdminHandler) AdminLogout(c *gin.Context) {
	c.SetCookie("AdminAuth", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "admin logouted",
		Data:       nil,
		Errors:     nil,
	})

}

func (cr *AdminHandler) BlockUser(c *gin.Context) {
	var body req.BlockData
	err := gin.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "binding fail",
			Data:       nil,
			Errors:     err,
		})
		return
	}

}
