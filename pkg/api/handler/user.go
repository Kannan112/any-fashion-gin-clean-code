package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	handlerUtil "github.com/kannan112/go-gin-clean-arch/pkg/api/handlerUril"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type UserHandler struct {
	userUseCase services.UserUseCase
	cartUseCase services.CartUseCases
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

//---------------------------------UserSignUp-----------------------------

// CreateAccount
// @Summary User Signup
// @ID UserSignUp
// @Description Login as a user to access the ecommerce site
// @Tags Users
// @Accept json
// @Produce json
// @Param   input   body     req.UserReq{}   true  "Input Field"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /user/signup [post]

func (cr *UserHandler) UserSignUp(ctx *gin.Context) {
	var user req.UserReq
	err := ctx.Bind(&user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, res.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userData, err := cr.userUseCase.UserSignUp(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "unable signup",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.cartUseCase.CreateCart(userData.Id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "unable create cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, res.Response{
		StatusCode: 200,
		Message:    "user signup Scart_itemsuccessfully",
		Data:       userData,
		Errors:     nil,
	})

}

// -------------------------------UserLogin-------------------

// LoginWithEmail
// @Summary User Login
// @ID UserLogin
// @Description Login as a user to access the ecommerce site
// @Tags Users
// @Accept json
// @Produce json
// @Param   input   body     req.LoginReq{}   true  "Input Field"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /user/login [post]
func (cr *UserHandler) UserLogin(c *gin.Context) {
	var user req.LoginReq
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to read request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	sessionValue, err := cr.userUseCase.UserLogin(c.Request.Context(), user)
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
	c.SetCookie("UserAuth", sessionValue, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "logined successfuly",
		Data:       nil,
		Errors:     nil,
	})

}

// -------------------------------UserLogout-------------------

// Logout
// @Summary User Logout
// @ID UserLogout
// @Description User logout to access the ecommerce site
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /user/logout [get]
func (cr *UserHandler) UserLogout(c *gin.Context) {
	c.SetCookie("UserAuth", "", 1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "UserLogouted",
	})
}
// AddAddress
// @Summary Add Address
// @ID AddAddress
// @Description Login as a user to access the ecommerce site
// @Tags Address
// @Accept json
// @Produce json
// @Param   input   body     req.Address{}   true  "Input Field"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /user/address/add [post]
func (cr *UserHandler) AddAddress(c *gin.Context) {
	id, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to get the userId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var address req.Address
	err = c.Bind(&address)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to bind Address",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.userUseCase.AddAddress(id, address)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Can't add address",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "address added successfully",
		Data:       address,
		Errors:     nil,
	})

}

// AddAddress
// @Summary Update Address
// @ID UpdateAddress
// @Description Login as a user to access the ecommerce site
// @Tags Address
// @Accept json
// @Produce json
// @Param   input   body     req.Address{}   true  "Input Field"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /user/address/update [patch]
func (cr *UserHandler) UpdateAddress(c *gin.Context) {
	paramsId := c.Param("addressId")
	addressId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Can't find ProductId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	id, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to get the userId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var address req.Address
	err = c.Bind(&address)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to bind Address",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.userUseCase.UpdateAddress(id, addressId, address)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Can't add address",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "address updated successfully",
		Data:       address,
		Errors:     nil,
	})
}

// AddAddress
// @Summary List Address
// @ID ListallAddres
// @Description Login as a user to access the ecommerce site
// @Tags Address
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /user/address/listall [get]
func (cr *UserHandler) ListallAddress(c *gin.Context) {
	id, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to get the userId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	data, err := cr.userUseCase.ListallAddress(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to list",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "All Address",
		Data:       data,
		Errors:     nil,
	})
}

// ViewProfile
// @Summary View Profile
// @ID ViewProfile
// @Description Login as a user to access the ecommerce site
// @Tags Profile
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /user/profile/view [get]
func (cr *UserHandler) ViewProfile(c *gin.Context) {
	id, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to find the id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	profile, err := cr.userUseCase.ViewProfile(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "Profile",
		Data:       profile,
		Errors:     nil,
	})
}

// EditProfile
// @Summary Edit Profile
// @ID EditProfile
// @Description Edit user prodile ecommerce site
// @Tags Profile
// @Accept json
// @Produce json
// @Param   input   body     req.UserReq{}   true  "Input Field"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /user/profile/edit [patch]
func (cr *UserHandler) EditProfile(c *gin.Context) {
	id, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusAccepted, res.Response{
			StatusCode: 400,
			Message:    "failed to find the id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var update req.UserReq
	profile, err := cr.userUseCase.EditProfile(id, update)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to update the profile",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "Updated successfully",
		Data:       profile,
		Errors:     nil,
	})
}
