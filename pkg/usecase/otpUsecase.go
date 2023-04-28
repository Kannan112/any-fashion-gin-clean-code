package usecase

import (
	"context"
	"os"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/config"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

type OtpUseCase struct {
	otpRepo interfaces.UserRepository
}

func NewOtpUseCase(cfg config.Config) services.OtpUseCase {
	return &OtpUseCase{
		
	}
}

func (c *OtpUseCase) SendOtp(ctx context.Context, phno req.OTPData) error {
	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTHTOKEN"),
	})

	params := &openapi.CreateVerificationParams{}
	params.SetTo(phno.PhoneNumber)
	params.SetChannel("sms")
	_, err := client.VerifyV2.CreateVerification(os.Getenv("TWILIO_SERVICES_ID"), params)
	return err
}

func (c *OtpUseCase) ValidateOtp(otpDetails req.VerifyOtp) (*openapi.VerifyV2VerificationCheck, error) {
	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTHTOKEN"),
	})
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(otpDetails.User.PhoneNumber)
	params.SetCode(otpDetails.Code)
	resp, err := client.VerifyV2.CreateVerificationCheck(os.Getenv("TWILIO_SERVICES_ID"), params)
	return resp, err
}
