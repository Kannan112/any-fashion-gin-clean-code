package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/config"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

type OtpUseCase struct {
	cfg      config.Config
	UserRepo interfaces.UserRepository
}

func NewOtpUseCase(cfg config.Config, UserRepo interfaces.UserRepository) services.OtpUseCase {
	return &OtpUseCase{
		cfg:      cfg,
		UserRepo: UserRepo,
	}
}

func (c *OtpUseCase) SendOtp(ctx context.Context, phno req.OTPData) error {
	check, err := c.UserRepo.IsSignIn(phno.PhoneNumber)
	if err != nil {
		return errors.New(err.Error())
	}
	if !check {
		return errors.New("failed to find your account")
	}
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: c.cfg.TWILIOACCOUNTSID,
		Password: c.cfg.TWILIOAUTHTOKEN,
	})
	params := &openapi.CreateVerificationParams{}
	params.SetTo(phno.PhoneNumber)
	params.SetChannel("sms")
	_, err = client.VerifyV2.CreateVerification(c.cfg.TWILIOSERVICESID, params)
	if err != nil {
		return err
	}
	return nil
}

func (c *OtpUseCase) VerifyOTP(ctx context.Context, userData req.Otpverifier) error {

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Password: c.cfg.TWILIOAUTHTOKEN,
		Username: c.cfg.TWILIOACCOUNTSID,
	})
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo(userData.Phone)
	params.SetCode(userData.Pin)
	resp, err := client.VerifyV2.CreateVerificationCheck(c.cfg.TWILIOSERVICESID, params)
	if err != nil {
		return errors.New("sorry, the otp has expired or is no longer valid. please generate a new otp to continue.")
	} else if *resp.Status == "approved" {
		fmt.Println("Correct!")
		err := c.UserRepo.AccountVerify(userData.Phone)
		if err != nil {
			return err
		}
		return nil
	} else {
		return fmt.Errorf("incorrect")
	}

}
