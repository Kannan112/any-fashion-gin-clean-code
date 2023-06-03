package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

type OtpUseCase interface {
	SendOtp(ctx context.Context, phone req.OTPData) (string, error)
	ValidateOtp(otpDetails req.VerifyOtp) (*openapi.VerifyV2VerificationCheck, error)
}
