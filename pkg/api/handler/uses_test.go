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

// ... (previous imports and code)
// package handler

// import (
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang/mock/gomock"
// 	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
// 	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
// 	mock_interfaces "github.com/kannan112/go-gin-clean-arch/pkg/usecase/mockusecase"
// 	"github.com/stretchr/testify/assert"
// )

// func TestUserSignUp(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	userUseCase := mock_interfaces.NewMockUserUseCase(ctrl)
// 	cartUseCase := mock_interfaces.NewMockCartUseCases(ctrl)
// 	walletUseCase := mock_interfaces.NewMockWalletUseCase(ctrl)

// 	UserHandler := NewUserHandler(userUseCase, cartUseCase, walletUseCase)

// 	testData := []struct {
// 		name             string
// 		userData         req.UserReq
// 		buildStub        func(userUsecase *mock_interfaces.MockUserUseCase, cartUsecase *mock_interfaces.MockCartUseCases)
// 		buildCartStub    func(cartUsecase *mock_interfaces.MockCartUseCases)
// 		buildWalletStub  func(walletUsecase *mock_interfaces.MockWalletUseCase)
// 		expectedCode     int
// 		expectedResponse res.Response
// 	}{
// 		{
// 			name: "new user",
// 			userData: req.UserReq{
// 				Name:     "akshay",
// 				Email:    "akshay@gmail.com",
// 				Mobile:   "9072001341",
// 				Password: "123456789",
// 			},
// 			buildStub: func(userUsecase *mock_interfaces.MockUserUseCase, cartUsecase *mock_interfaces.MockCartUseCases) {
// 				userUsecase.EXPECT().UserSignUp(gomock.Any(), req.UserReq{
// 					Name:     "akshay",
// 					Email:    "akshay@gmail.com",
// 					Mobile:   "9072001341",
// 					Password: "123456789",
// 				}).Times(1).
// 					Return(res.UserData{
// 						Id:     1,
// 						Name:   "akshay",
// 						Email:  "akshay@gmail.com",
// 						Mobile: "9072001341",
// 					}, nil)
// 				cartUsecase.EXPECT().CreateCart(1).Times(1).Return(nil)
// 			},
// 			buildCartStub: func(cartUsecase *mock_interfaces.MockCartUseCases) {
// 				cartUsecase.EXPECT().CreateCart(1).Times(1).Return(nil)
// 			},
// 			buildWalletStub: func(walletUsecase *mock_interfaces.MockWalletUseCase) {
// 				walletUsecase.EXPECT().SaveWallet(gomock.Any(), 1).Times(1).Return(nil)
// 			},
// 			expectedCode: 200,
// 			expectedResponse: res.Response{
// 				StatusCode: 200,
// 				Message:    "user signup successfully",
// 				Data: res.UserData{
// 					Id:     1,
// 					Name:   "akshay",
// 					Email:  "akshay@gmail.com",
// 					Mobile: "9072001341",
// 				},
// 				Errors: nil,
// 			},
// 		},
// 		{
// 			name: "duplicate user",
// 			userData: req.UserReq{
// 				Name:     "akshay",
// 				Email:    "akshay@gmail.com",
// 				Mobile:   "9072001341",
// 				Password: "123456789",
// 			},
// 			buildStub: func(userUsecase *mock_interfaces.MockUserUseCase, cartUsecase *mock_interfaces.MockCartUseCases) {
// 				userUsecase.EXPECT().UserSignUp(gomock.Any(), req.UserReq{
// 					Name:     "akshay",
// 					Email:    "akshay@gmail.com",
// 					Mobile:   "9072001341",
// 					Password: "123456789",
// 				}).Times(1).
// 					Return(
// 						res.UserData{},
// 						errors.New("user already exists"),
// 					)
// 			},
// 			buildCartStub: func(cartUsecase *mock_interfaces.MockCartUseCases) {
// 				// This scenario doesn't test the cartUseCase
// 				// No expectations are set for cartUseCase
// 			},
// 			buildWalletStub: func(walletUsecase *mock_interfaces.MockWalletUseCase) {
// 				// This scenario doesn't test the walletUseCase
// 				// No expectations are set for walletUseCase
// 			},
// 			expectedCode: 400,
// 			expectedResponse: res.Response{
// 				StatusCode: 400,
// 				Message:    "unable signup",
// 				Data:       res.UserData{},
// 				Errors:     "user already exists",
// 			},
// 		},
// 		{
// 			name:     "invalid request",
// 			userData: req.UserReq{
// 				// This user data is intentionally left empty to trigger binding failure
// 			},
// 			buildStub: func(userUsecase *mock_interfaces.MockUserUseCase, cartUsecase *mock_interfaces.MockCartUseCases) {
// 				// This scenario doesn't test the userUseCase
// 				// No expectations are set for userUseCase
// 			},
// 			buildCartStub: func(cartUsecase *mock_interfaces.MockCartUseCases) {
// 				// This scenario doesn't test the cartUseCase
// 				// No expectations are set for cartUseCase
// 			},
// 			buildWalletStub: func(walletUsecase *mock_interfaces.MockWalletUseCase) {
// 				// This scenario doesn't test the walletUseCase
// 				// No expectations are set for walletUseCase
// 			},
// 			expectedCode: 422,
// 			expectedResponse: res.Response{
// 				StatusCode: 422,
// 				Message:    "can't bind",
// 				Data:       nil,
// 				Errors:     "EOF", // This will be the error message when binding fails
// 			},
// 		},
// 	}

// 	for _, tt := range testData {
// 		t.Run(tt.name, func(t *testing.T) {
// 			userUsecase := mock_interfaces.NewMockUserUseCase(ctrl)
// 			cartUsecase := mock_interfaces.NewMockCartUseCases(ctrl)
// 			walletUsecase := mock_interfaces.NewMockWalletUseCase(ctrl)

// 			tt.buildStub(userUsecase, cartUsecase)
// 			tt.buildCartStub(cartUsecase)
// 			tt.buildWalletStub(walletUsecase)

// 			//	UserHandler := NewUserHandler(userUsecase, cartUsecase, walletUsecase)

// 			gin.SetMode(gin.TestMode)
// 			engine := gin.Default()
// 			engine.POST("/user/signup", UserHandler.UserSignUp)

// 			// Marshal an empty JSON to simulate binding failure
// 			body, err := json.Marshal(tt.userData)
// 			assert.NoError(t, err)

// 			url := "/user/signup"
// 			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
// 			recorder := httptest.NewRecorder()

// 			engine.ServeHTTP(recorder, req)

// 			var actual res.Response
// 			err = json.Unmarshal(recorder.Body.Bytes(), &actual)
// 			assert.NoError(t, err)

// 			assert.Equal(t, tt.expectedCode, recorder.Code)
// 			assert.Equal(t, tt.expectedResponse.Message, actual.Message)
// 			assert.Equal(t, tt.expectedResponse.Data, actual.Data)
// 			assert.Equal(t, tt.expectedResponse.Errors, actual.Errors)
// 		})
// 	}
// }
