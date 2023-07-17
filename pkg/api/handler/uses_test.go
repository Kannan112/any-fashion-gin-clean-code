package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	mock_interfaces "github.com/kannan112/go-gin-clean-arch/pkg/usecase/mockusecase"
	"github.com/stretchr/testify/assert"
)

func TestUserSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	userUseCase := mock_interfaces.NewMockUserUseCase(ctrl)
	cartUseCase := mock_interfaces.NewMockCartUseCases(ctrl)
	walletUseCase := mock_interfaces.NewMockWalletUseCase(ctrl)
	UserHandler := NewUserHandler(userUseCase, cartUseCase, walletUseCase)

	testData := []struct {
		name             string
		userData         req.UserReq
		buildStub        func(userUsecase mock_interfaces.MockUserUseCase)
		expectedCode     int
		expectedResponse res.Response
		expectedData     res.UserData
		expectedError    error
	}{
		{
			name: "new user",
			userData: req.UserReq{
				Name:     "abhinand",
				Email:    "abhinand@gmail.com",
				Mobile:   "9072001341",
				Password: "123456789",
			},
			buildStub: func(userUsecase mock_interfaces.MockUserUseCase) {
				userUsecase.EXPECT().UserSignUp(gomock.Any(), req.UserReq{
					Name:     "abhinand",
					Email:    "abhinand@gmail.com",
					Mobile:   "9072001341",
					Password: "123456789",
				}).Times(1).
					Return(res.UserData{
						Id:     1,
						Name:   "abhinand",
						Email:  "abhinand@gmail.com",
						Mobile: "9072001341",
					}, nil)
				cartUseCase.EXPECT().CreateCart(1).Times(1).Return(nil)
				walletUseCase.EXPECT().SaveWallet(gomock.Any(), 1).Times(1).Return(nil)
			},
			expectedCode: 200,
			expectedResponse: res.Response{
				StatusCode: 200,
				Message:    "user signup successfully",
				Data: res.UserData{
					Id:     1,
					Name:   "abhinand",
					Email:  "abhinand@gmail.com",
					Mobile: "9072001341",
				},
				Errors: nil,
			},
			expectedData: res.UserData{
				Id:     1,
				Name:   "abhinand",
				Email:  "abhinand@gmail.com",
				Mobile: "9072001341",
			},
			expectedError: nil,
		},
		{
			name: "duplicate user",
			userData: req.UserReq{
				Name:     "abhinand",
				Email:    "abhinand@gmail.com",
				Mobile:   "9072001341",
				Password: "123456789",
			},
			buildStub: func(userUsecase mock_interfaces.MockUserUseCase) {
				userUsecase.EXPECT().UserSignUp(gomock.Any(), req.UserReq{
					Name:     "abhinand",
					Email:    "abhinand@gmail.com",
					Mobile:   "9072001341",
					Password: "123456789",
				}).Times(1).
					Return(
						res.UserData{},
						errors.New("user already exists"),
					)
				// cartUseCase.EXPECT().CreateCart(1).Times(1).Return(nil)
			},
			expectedCode: 400,
			expectedResponse: res.Response{
				StatusCode: 400,
				Message:    "unable signup",
				Data:       res.UserData{},
				Errors:     "user already exits",
			},
			expectedData:  res.UserData{},
			expectedError: errors.New("user already exists"),
		},
	}
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub(*userUseCase)
			engine := gin.Default()            //create an engin instance
			recorder := httptest.NewRecorder() //creeate a responce recorder to capture the responce from the request
			engine.POST("/user/signup", UserHandler.UserSignUp)
			var body []byte
			fmt.Println(tt.userData)
			body, err := json.Marshal(tt.userData) //marshal the user data field into json
			assert.NoError(t, err)
			url := "/user/signup"

			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body)) //create a new http request
			engine.ServeHTTP(recorder, req)                                         //execute the http req
			var actual res.Response
			err = json.Unmarshal(recorder.Body.Bytes(), &actual) //unmarshal the op
			assert.NoError(t, err)
			// fmt.Println(actual.StatusCode)
			assert.Equal(t, tt.expectedCode, actual.StatusCode)
			assert.Equal(t, tt.expectedResponse.Message, actual.Message)

		})
	}
}
