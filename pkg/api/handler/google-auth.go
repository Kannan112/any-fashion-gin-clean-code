package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	"github.com/kannan112/go-gin-clean-arch/pkg/config"
	"github.com/kannan112/go-gin-clean-arch/pkg/domain"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
	"github.com/markbates/goth/providers/google"
	"golang.org/x/oauth2"
)

type AuthHandler struct {
	UserAuthUsecase  services.UserUseCase
	AdminAuthUsecase services.AdminUsecase
	AuthUseCase      services.AuthUserCase
	Config           config.Config
}

func NewAuthHandler(userAuthUsecase services.UserUseCase, adminAuthUsecase services.AdminUsecase, config config.Config, AuthUseCase services.AuthUserCase) *AuthHandler {
	return &AuthHandler{
		UserAuthUsecase:  userAuthUsecase,
		AdminAuthUsecase: adminAuthUsecase,
		Config:           config,
		AuthUseCase:      AuthUseCase,
	}
}

func SetUpConfig(c *AuthHandler) *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     c.Config.GoauthClientID,
		ClientSecret: c.Config.GoauthClientSRC,
		RedirectURL:  c.Config.Redirect_URL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
	return conf
}

// @Summary Google Login
// @Description Initiates the Google login flow
// @Tags Authentication
// @Produce json
// @Success 303 {string} string "See Other"
// @Router /api/auth/google-login [get]
func (auth *AuthHandler) GoogleLogin(ctx *gin.Context) {
	googleConfig := SetUpConfig(auth) // &c.Config
	url := googleConfig.AuthCodeURL("randomstate")
	ctx.Redirect(http.StatusSeeOther, url)
}

// @Summary Google Auth Callback
// @Description Callback endpoint after Google authentication
// @Tags Authentication
// @Accept json
// @Produce json
// @Param code query string true "Authorization Code"
// @Param state query string true "State"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /api/auth/google-callback [get]
func (c *AuthHandler) GoogleAuthCallback(ctx *gin.Context) {
	state := ctx.Query("state")
	if state != "randomstate" {
		ctx.JSON(http.StatusNonAuthoritativeInfo, res.Response{
			StatusCode: http.StatusNonAuthoritativeInfo,
			Message:    "some internal server error occured",
			Data:       nil,
			Errors:     "failed to match with state",
		})
		return
	}
	code := ctx.Query("code")
	googleConfig := SetUpConfig(c)
	token, err := googleConfig.Exchange(ctx, code)
	if err != nil {
		ctx.JSON(http.StatusNonAuthoritativeInfo, res.Response{
			StatusCode: http.StatusNonAuthoritativeInfo,
			Message:    "Error exchanging code for token",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)

	if err != nil {
		ctx.JSON(http.StatusNonAuthoritativeInfo, res.Response{
			StatusCode: http.StatusNonAuthoritativeInfo,
			Message:    "Error getting user info from Google:",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	authData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.JSON(http.StatusNonAuthoritativeInfo, res.Response{
			StatusCode: http.StatusNonAuthoritativeInfo,
			Message:    "failed to read:",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	var userInfo domain.GoAuthUserInfo
	err = json.Unmarshal(authData, &userInfo)
	if err != nil {
		ctx.JSON(http.StatusNonAuthoritativeInfo, res.Response{
			StatusCode: http.StatusNonAuthoritativeInfo,
			Message:    "failed to write:",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	fmt.Printf("ID: %s\nEmail: %s\nName: %s\n", userInfo.ID, userInfo.Email, userInfo.Name)

	data := req.GoogleAuth{
		Name:  userInfo.Name,
		Email: userInfo.Email,
	}
	fmt.Println("testingAneERror")
	accessToken, refreshToken, err := c.AuthUseCase.GoogleLoginUser(ctx, data)
	if err != nil {
		ctx.JSON(http.StatusNonAuthoritativeInfo, res.Response{
			StatusCode: http.StatusBadGateway,
			Message:    "some internal server error occured",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	tokens := map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}

	fmt.Println("testingAne")

	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: http.StatusOK,
		Message:    "successfully login",
		Data:       tokens,
		Errors:     nil,
	})

}
