package usecase

import (
	"context"
	"fmt"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/config"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

type OtpUseCase struct {
	cfg config.Config
}

func NewOtpUseCase(cfg config.Config) services.OtpUseCase {
	return &OtpUseCase{
		cfg: cfg,
	}
}

func (c *OtpUseCase) SendOtp(ctx context.Context, phno req.OTPData) error {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: c.cfg.TWILIOACCOUNTSID,
		Password: c.cfg.TWILIOAUTHTOKEN,
	})
	params := &openapi.CreateVerificationParams{}
	params.SetTo(phno.PhoneNumber)
	params.SetChannel("sms")
	_, err := client.VerifyV2.CreateVerification(c.cfg.TWILIOSERVICESID, params)
	return err
}

//	func (c *OtpUseCase) ValidateOtp(otpDetails req.VerifyOtp) (*openapi.VerifyV2VerificationCheck, error) {
//		var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
//			Username: c.cfg.TWILIOACCOUNTSID,
//			Password: c.cfg.TWILIOAUTHTOKEN,
//		})
//		params := &openapi.CreateVerificationCheckParams{}
//		params.SetTo(otpDetails.User.PhoneNumber)
//		params.SetCode(otpDetails.Code)
//		resp, err := client.VerifyV2.CreateVerificationCheck(c.cfg.TWILIOSERVICESID, params)
//		return resp, err
//	}
func (c *OtpUseCase) VerifyOTP(ctx context.Context, userData req.Otpverifier) error {

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Password: c.cfg.TWILIOAUTHTOKEN,
		Username: c.cfg.TWILIOACCOUNTSID,
	})
	fmt.Println("phone", userData.Phone, "otp", userData.Pin)
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(userData.Phone)
	params.SetCode(userData.Pin)
	resp, err := client.VerifyV2.CreateVerificationCheck(c.cfg.TWILIOSERVICESID, params)
	if err != nil {
		return fmt.Errorf(err.Error())
	} else if *resp.Status == "approved" {
		fmt.Println("Correct!")
		return nil
	} else {
		return fmt.Errorf("incorrect")
	}
	return nil
}
