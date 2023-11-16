package interfaces

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

type AuthUserCase interface {
	//
	GoogleLoginUser(ctx context.Context, googleuser domain.Users) (string, string, error)
	/*
		pending :
				AdminUseCase TO AuthUseCase -> admin login,admin signup
				UserUseCase TO AuthUseCase -> userLogin,userSignup....
	*/
}
