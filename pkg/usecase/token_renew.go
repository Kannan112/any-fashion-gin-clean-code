package usecase

import (
	"context"
	"fmt"

	"github.com/kannan112/go-gin-clean-arch/pkg/api/middleware/token"
	interfaces "github.com/kannan112/go-gin-clean-arch/pkg/repository/interface"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type TokenRenewUseCase struct {
	token interfaces.RefreshTokenRepository
}

func NewTokenRenewUseCase(tokenRepo interfaces.RefreshTokenRepository) services.RenewTokenUseCase {
	return &TokenRenewUseCase{
		token: tokenRepo,
	}
}

func (c *TokenRenewUseCase) GetAccessToken(ctx context.Context, AccessToken string) (string, error) {
	claims, err := token.RefreshTokenClaims(AccessToken)
	if err != nil {
		return "failed to claim data", err
	}
	switch claims.Role {
	case "admin":
		RefrshToken, err := c.token.AdminFindRefreshToken(ctx, claims.ID)
		if err != nil {
			return "", fmt.Errorf("failed to find refresh token admin side got an error%d", err)
		}
		claims, err := token.RefreshTokenClaims(RefrshToken)
		if err != nil {
			return "", err
		}
		NewAccessToken, err := token.GenerateAccessToken(int(claims.ID), claims.Role)
		if err != nil {
			return "", fmt.Errorf("failed to generate access token")
		}
		return NewAccessToken, nil

	case "user":
		RefrshToken, err := c.token.UserFindRefreshToken(ctx, claims.ID)
		if err != nil {
			return "", fmt.Errorf("failed to find refresh token user side got an error%d", err)
		}
		claims, err := token.RefreshTokenClaims(RefrshToken)
		if err != nil {
			return "", err
		}
		NewAccessToken, err := token.GenerateAccessToken(int(claims.ID), claims.Role)
		if err != nil {
			return "", fmt.Errorf("failed to generate access token")
		}
		return NewAccessToken, nil

	case "":

		return "", fmt.Errorf("cant find any role in claims")
	}

	return "nothing left", nil
}
