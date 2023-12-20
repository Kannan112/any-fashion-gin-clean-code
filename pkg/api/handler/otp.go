package handler

import (
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
// @Router /api/user/otp/send [post]
func (cr *OtpHandler) SendOtp(c *gin.Context) {
	var phno req.OTPData
	err := c.Bind(&phno)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 422,
			Message:    "unable to process the request",
			Data:       nil,
			Errors:     err.Error(),
		})
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

	if !isSignIn {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "no user found",
			Data:       nil,
			Errors:     nil,
		})
		return
	}
	err = cr.otpUseCase.SendOtp(c.Request.Context(), phno)
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
		Data:       phno.PhoneNumber,
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
// @Param otp body req.Otpverifier true "OTP sent to user's mobile number"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /api/user/otp/verify [post]
func (cr *OtpHandler) ValidateOtp(c *gin.Context) {
	var otpDetails req.Otpverifier
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
	err = cr.otpUseCase.VerifyOTP(c, otpDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "validation failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, res.Response{
		StatusCode: 201,
		Message:    "user account verified",
		Data:       otpDetails.Phone,
		Errors:     nil,
	})
}
