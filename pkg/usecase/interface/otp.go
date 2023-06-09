package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
)

type OtpUseCase interface {
	SendOtp(ctx context.Context, phone req.OTPData) error
	VerifyOTP(ctx context.Context, userData req.Otpverifier) error
}
