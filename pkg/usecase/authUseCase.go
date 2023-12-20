package usecase

import (
	"context"
	"errors"

	"github.com/kannan112/go-gin-clean-arch/pkg/api/middleware/token"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"

	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
)

type AuthUseCase struct {
	UserRepo  interfaces.UserRepository
	TokenRepo interfaces.RefreshTokenRepository
}

func NewAuthUseCase(Repo interfaces.UserRepository, token interfaces.RefreshTokenRepository) services.AuthUserCase {
	return &AuthUseCase{
		UserRepo:  Repo,
		TokenRepo: token,
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
		AccessTokenString, err := token.GenerateAccessToken(int(data.Id), "user")
		if err != nil {
			return "", "", err
		}
		RefreshTokenString, err := token.GenerateRefreshToken(int(data.Id), "user")
		if err != nil {
			return "", "", err
		}
		return AccessTokenString, RefreshTokenString, nil

	}

	AccessTokenString, err := token.GenerateAccessToken(int(clean.ID), "user")
	if err != nil {
		return "", "", err
	}
	RefreshTokenString, err := token.GenerateRefreshToken(int(clean.ID), "user")
	if err != nil {
		return "", "", err
	}
	if err := c.TokenRepo.UserRefreshTokenAdd(RefreshTokenString, clean.ID); err != nil {
		return "", "", errors.New("failed to save refresh token")
	}

	return AccessTokenString, RefreshTokenString, nil
}
