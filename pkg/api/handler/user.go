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
	userUseCase   services.UserUseCase
	cartUseCase   services.CartUseCases
	walletUseCase services.WalletUseCase
}

func NewUserHandler(usecase services.UserUseCase, cartcase services.CartUseCases, walletUseCase services.WalletUseCase) *UserHandler {
	return &UserHandler{
		userUseCase:   usecase,
		cartUseCase:   cartcase,
		walletUseCase: walletUseCase,
	}
}

//---------------------------------UserSignUp-----------------------------

// UserSignUp handles user signup.
// @Summary user-signup
// @ID UserSignUp
// @Description Signup as a new user to access the ecommerce site
// @Tags Users
// @Accept json
// @Produce json
// @Param user_details body  req.UserReq true "User details"
// @Success 201 {object} res.Response
// @Failure 400 {object} res.Response
// @Security BearerTokenAuth
// @Router /api/user/signup [post]
func (cr *UserHandler) UserSignUp(ctx *gin.Context) {
	var user req.UserReq
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, res.Response{
			StatusCode: 400,
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
	err = cr.walletUseCase.SaveWallet(ctx, userData.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to create wallet",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res.Response{
		StatusCode: 200,
		Message:    "user signup successfully",
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
// @Security BearerTokenAuth
// @Router /api/user/login [post]
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
	Token, err := cr.userUseCase.UserLogin(c.Request.Context(), user)
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
		Message:    "logined successfuly",
		Data:       Token,
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
// @Security BearerTokenAuth
// @Router /api/user/logout [get]
func (cr *UserHandler) UserLogout(c *gin.Context) {
	c.SetCookie("UserAuth", "", 1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "UserLogouted",
	})
}

// AddAddress
// @Summary Add Address
// @ID add-address
// @Description Login as a user to access the ecommerce site
// @Tags Address
// @Accept json
// @Produce json
// @Param   input   body     req.AddAddress{}   true  "Input Field"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Security BearerTokenAuth
// @Router /api/user/address/add [post]
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
	var address req.AddAddress
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

// UpdateAddress
// @Summary Update Address
// @ID update-address
// @Description user update addresses
// @Tags Address
// @Accept json
// @Produce json
// @Param addressId path string true "addressId"
// @Param   input   body     req.AddAddress{}   true  "Input Field"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Security BearerTokenAuth
// @Router /api/user/address/update/{addressId} [patch]
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
	var address req.AddAddress
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
			Message:    "failed to update",
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

// ListAddress
// @Summary List Addresses
// @ID list-all-addresses
// @Description Login as a user to access the ecommerce site
// @Tags Address
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Security BearerTokenAuth
// @Router /api/user/address/list [get]
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
// @Security BearerTokenAuth
// @Router /api/user/profile/view [get]
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
// @Security BearerTokenAuth
// @Router /api/user/profile/edit [patch]
func (cr *UserHandler) EditProfile(c *gin.Context) {
	id, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusAccepted, res.Response{
			StatusCode: 400,
			Message:    "User not login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var update req.UserReq
	err = c.Bind(&update)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed binding",
			Data:       nil,
			Errors:     err,
		})
		return
	}
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

// DeleteAddress
// @Summary Delete Address
// @ID delete-address
// @Description user can delete any of his addresses
// @Tags Address
// @Accept json
// @Produce json
// @Param addressId path string true "addressId"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Security BearerTokenAuth
// @Router /api/user/address/delete/{addressId} [delete]
func (cr *UserHandler) DeleteAddress(ctx *gin.Context) {
	addressStr := ctx.Param("addressId")
	addressId, err := strconv.Atoi(addressStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to catch id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userId, err := handlerUtil.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "User not login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	domain, err := cr.userUseCase.DeleteAddress(ctx, userId, addressId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "Address Deleted",
		Data:       domain,
		Errors:     nil,
	})

}
