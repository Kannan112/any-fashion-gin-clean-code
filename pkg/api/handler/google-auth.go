package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/config"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type AuthHandler struct {
	UserAuthUsecase  services.UserUseCase
	AdminAuthUsecase services.AdminUsecase
	AuthUseCase      services.AuthUserCase
	Config           config.Config
}

func NewAuthHandler(userAuthUsecase services.UserUseCase, adminAuthUsecase services.AdminUsecase, config config.Config) *AuthHandler {
	return &AuthHandler{
		UserAuthUsecase:  userAuthUsecase,
		AdminAuthUsecase: adminAuthUsecase,
		Config:           config,
	}
}

func (c *AuthHandler) UseGoogleAuthLoginPage(ctx *gin.Context) {
	ctx.HTML(200, "googleauth.html", nil)
}

func (c *AuthHandler) UserGoogleAuthInitialize(ctx *gin.Context) {
	goauthclientID := c.Config.GoauthClientID
	goauthClientSrc := c.Config.GoauthClientSRC
	callbackUrl := c.Config.Redirect_URL
	goth.UseProviders(
		google.New(goauthclientID, goauthClientSrc, callbackUrl, "email", "profile"),
	)

	// start google login
	gothic.BeginAuthHandler(ctx.Writer, ctx.Request)
}

func (c *AuthHandler) UserGoogoleAuthCallBack(ctx *gin.Context) {
	googleuser, err := gothic.CompleteUserAuth(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, res.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "failed to get user details",
			Data:       nil,
			Errors:     err,
		})
	}
	data := domain.Users{
		Name:   googleuser.Name,
		Email:  googleuser.Email,
		Images: googleuser.AvatarURL,
	}

	accessToken, refreshToken, err := c.AuthUseCase.GoogleLoginUser(ctx, data)
	tokens := map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}
	if err != nil {
		ctx.JSON(http.StatusNonAuthoritativeInfo, res.Response{
			StatusCode: http.StatusNonAuthoritativeInfo,
			Message:    "some internal server error occured",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: http.StatusOK,
		Message:    "successfully login",
		Data:       tokens,
		Errors:     nil,
	})

}
