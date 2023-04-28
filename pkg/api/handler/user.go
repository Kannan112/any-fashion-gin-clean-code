package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	handlerUtil "github.com/kannan112/go-gin-clean-arch/pkg/api/handlerUril"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

//---------------------------------_UserSignUp-----------------------------

func (cr *UserHandler) UserSignUp(c *gin.Context) {
	var user req.UserReq
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, res.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	userData, err := cr.userUseCase.UserSignUp(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "unable signup",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	c.JSON(http.StatusCreated, res.Response{
		StatusCode: 201,
		Message:    "user signup Successfully",
		Data:       userData,
		Errors:     nil,
	})

}

//-------------------------------UserLogin-------------------

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
func (cr *UserHandler) UserLogout(c *gin.Context) {
	c.SetCookie("UserAuth", "", 1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "UserLogouted",
	})
}
func (cr *UserHandler) AddAddress(c *gin.Context) error {
	id, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to get the userId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return err
	}
	var address req.Address
	err = c.Bind(address)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to bind Address",
			Data:       nil,
			Errors:     err.Error(),
		})
		return err
	}
	err = cr.userUseCase.AddAddress(id, address)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Can't add address",
			Data:       nil,
			Errors:     err.Error(),
		})
		
	}
	return fmt.Errorf("")

}
