package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	handlerUtil "github.com/kannan112/go-gin-clean-arch/pkg/api/handlerUril"
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
// @Param admin body req.CreateAdmin true "New Admin details"
// @Success 200 {object} res.Response
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

// @Summary Admin Login
// @Description Logs in an admin user and returns an authentication token
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin body req.LoginReq true "Admin login details"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /admin/login [post]
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

// @AdminLogout
// @Summary Admin Logout
// @ID AdminLogout
// @Description Logout the currently authenticated admin user
// @Tags admin
// @Produce json
// @Success 200 {object} res.Respons
// @Failure 400 {object} res.Respons
// @Router /admin/logout [get]

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
	errx := cr.adminUseCase.BlockUser(body, adminId)
	if errx != nil {
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
	err := ctx.Bind(&userEmail)
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
