package usecase

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/api/middleware/token"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
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

func (c *AuthUseCase) GoogleLoginUser(ctx context.Context, googleuser domain.Users) (string, string, error) {
	data, err := c.UserRepo.UserLogin(ctx, googleuser.Email)
	if err != nil {
		return "", "", err
	}
	AccessTokenString, err := token.JWTAccessTokenGen(int(data.ID), "user")
	if err != nil {
		return "", "", err
	}
	RefreshTokenString, err := token.JWTRefreshTokenGen(int(data.ID), "user")
	if err != nil {
		return "", "", err
	}
	return AccessTokenString, RefreshTokenString, nil
}
