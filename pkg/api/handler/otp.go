package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/config"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type OtpHandler struct {
	otpUseCase  services.OtpUseCase
	userUseCase services.UserUseCase
	cfg         config.Config
}

func NewOtpHandler(cfg config.Config, otpUseCase services.OtpUseCase, userUseCase services.UserUseCase) *OtpHandler {
	return &OtpHandler{
		otpUseCase:  otpUseCase,
		userUseCase: userUseCase,
		cfg:         cfg,
	}
}

// SendOtp
// @Summary Send OTP to user's mobile
// @ID send-otp
// @Description Send OTP to use's mobile
// @Tags Otp
// @Accept json
// @Produce json
// @Param user_mobile body  req.OTPData true "User mobile number"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /user/otp/send [post]
func (cr *OtpHandler) SendOtp(c *gin.Context) {
	var phno req.OTPData
	err := c.Bind(&phno)
	if err != nil {
		fmt.Println("e1")
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 422,
			Message:    "unable to process the request",
			Data:       nil,
			Errors:     err.Error(),
		})
		fmt.Println("e2")
		return
	}

	isSignIn, err := cr.userUseCase.IsSignIn(phno.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "No user with this phonenumber",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	fmt.Println(isSignIn)

	if !isSignIn {
		fmt.Println("login err")
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "no user found",
			Data:       nil,
			Errors:     nil,
		})
		fmt.Println("login err2")
		return
	}
	sid, err := cr.otpUseCase.SendOtp(c, phno)

	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "creatingfailed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, res.Response{
		StatusCode: 201,
		Message:    "otp send",
		Data:       sid,
		Errors:     nil,
	})
}

// ValidateOtp
// @Summary Validate the OTP to user's mobile
// @ID validate-otp
// @Description Validate the  OTP sent to use's mobile
// @Tags Otp
// @Accept json
// @Produce json
// @Param otp body req.VerifyOtp true "OTP sent to user's mobile number"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /user/otp/verify [post]
func (cr *OtpHandler) ValidateOtp(c *gin.Context) {
	var otpDetails req.VerifyOtp
	err := c.Bind(&otpDetails)

	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	resp, err := cr.otpUseCase.ValidateOtp(otpDetails)

	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
	} else if *resp.Status != "approved" {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "incorect",
			Data:       nil,
			Errors:     "incorect",
		})
		return
	}
	ss, err := cr.userUseCase.OtpLogin(otpDetails.User.PhoneNumber)
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
	c.SetCookie("UserAuth", ss, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "login successful",
		Data:       nil,
		Errors:     nil,
	})
}
