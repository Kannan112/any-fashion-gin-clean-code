package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	handlerUtil "github.com/kannan112/go-gin-clean-arch/pkg/api/handlerUril"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type WalletHandler struct {
	walletUseCase services.WalletUseCase
}

func NewWalletHandler(WalletUseCase services.WalletUseCase) *WalletHandler {
	return &WalletHandler{
		walletUseCase: WalletUseCase,
	}
}

// WalletProfile
// @Summary Wallet Profile
// @ID wallet-profile
// @Description user wallet displays super-coins
// @Tags Profile Wallet
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /api/user/wallet [get]
func (c *WalletHandler) WallerProfile(ctx *gin.Context) {
	userid, err := handlerUtil.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "please login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	profile, err := c.walletUseCase.GetWalletProfile(ctx, uint(userid))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to dispaly profile",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "user wallet profile",
		Data:       profile,
		Errors:     nil,
	})

}

// ApplyWallet
// @Summary Wallet Profile
// @ID apply-wallet
// @Description user can add wallet coins to the cart
// @Tags Profile Wallet
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /api/user/wallet/apply [post]
func (c *WalletHandler) ApplyWallet(ctx *gin.Context) {
	userid, err := handlerUtil.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "please login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = c.walletUseCase.ApplyWallet(ctx, uint(userid))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to apply wallet",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "coin added",
		Data:       nil,
		Errors:     nil,
	})
}

// Remove Wallet
// @Summary Wallet Profile
// @ID remove-wallet
// @Description user can add wallet coins to the cart
// @Tags Profile Wallet
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /api/user/wallet/remove [patch]
func (c *WalletHandler) RemoveWallet(ctx *gin.Context) {
	userid, err := handlerUtil.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "please login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = c.walletUseCase.RemoveWallet(ctx, uint(userid))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to remove wallet",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "coin removed",
		Data:       nil,
		Errors:     nil,
	})

}
