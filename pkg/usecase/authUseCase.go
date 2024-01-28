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
	UserRepo   interfaces.UserRepository
	TokenRepo  interfaces.RefreshTokenRepository
	CartRepo   interfaces.CartRepository
	walletRepo interfaces.WalletRepo
}

func NewAuthUseCase(Repo interfaces.UserRepository, token interfaces.RefreshTokenRepository, cart interfaces.CartRepository, wallet interfaces.WalletRepo) services.AuthUserCase {
	return &AuthUseCase{
		UserRepo:   Repo,
		TokenRepo:  token,
		CartRepo:   cart,
		walletRepo: wallet,
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
		if err := c.CartRepo.CreateCart(int(data.Id)); err != nil {
			return "", "", errors.New("failed to create cart")
		}
		err = c.walletRepo.SaveWallet(ctx, int(data.Id))
		if err != nil {
			return "", "", errors.New("failed to create user wallet")
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
