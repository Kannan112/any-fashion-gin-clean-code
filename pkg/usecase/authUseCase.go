package usecase

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/api/middleware/token"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"

	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
)

type AuthUseCase struct {
	UserRepo interfaces.UserRepository
}

func NewAuthUseCase(Repo interfaces.UserRepository) services.AuthUserCase {
	return &AuthUseCase{
		UserRepo: Repo,
	}
}

func (c *AuthUseCase) GoogleLoginUser(ctx context.Context, googleuser req.GoogleAuth) (string, string, error) {

	clean, err := c.UserRepo.UserLogin(ctx, googleuser.Email)
	if err != nil {
		return "", "", err
	}
	if clean.ID == 0 {
		data, err := c.UserRepo.AuthSignUp(googleuser)
		if err != nil {
			return "", "", err
		}
		AccessTokenString, err := token.JWTAccessTokenGen(int(data.Id), "user")
		if err != nil {
			return "", "", err
		}
		RefreshTokenString, err := token.JWTRefreshTokenGen(int(data.Id), "user")
		if err != nil {
			return "", "", err
		}
		return AccessTokenString, RefreshTokenString, nil

	}

	AccessTokenString, err := token.JWTAccessTokenGen(int(clean.ID), "user")
	if err != nil {
		return "", "", err
	}
	RefreshTokenString, err := token.JWTRefreshTokenGen(int(clean.ID), "user")
	if err != nil {
		return "", "", err
	}
	return AccessTokenString, RefreshTokenString, nil
}
