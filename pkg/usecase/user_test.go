package usecase

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	mockRepo "github.com/kannan112/go-gin-clean-arch/pkg/repository/mockrepo"
	"golang.org/x/crypto/bcrypt"
)

type eqCreateUserParamsMatcher struct {
	arg      req.UserReq
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(req.UserReq)
	if !ok {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(arg.Password), []byte(e.password)); err != nil {
		return false
	}
	e.arg.Password = arg.Password
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg req.UserReq, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func TestUserSignup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepo := mockRepo.NewMockUserRepository(ctrl)
	TokenRepo := mockRepo.NewMockRefreshTokenRepository(ctrl)
	userUseCase := NewUserUseCase(userRepo, TokenRepo)
	testData := []struct {
		name           string
		input          req.UserReq
		buildStub      func(userRepo mockRepo.MockUserRepository)
		expectedOutput res.UserData
		expectedError  error
	}{
		{
			name: "new user",
			input: req.UserReq{
				Name:     "abhinand",
				Email:    "abhi@gmail.com",
				Mobile:   "9072001341",
				Password: "123456789",
			},
			buildStub: func(userRepo mockRepo.MockUserRepository) {
				userRepo.EXPECT().UserSignUp(gomock.Any(),
					EqCreateUserParams(req.UserReq{
						Name:     "abhinand",
						Email:    "abhi@gmail.com",
						Mobile:   "9072001341",
						Password: "123456789",
					},
						"123456789")).
					Times(1).
					Return(res.UserData{
						Id:     1,
						Name:   "abhinand",
						Email:  "abhi@gmail.com",
						Mobile: "9072001341",
					}, nil)
			},
			expectedOutput: res.UserData{
				Id:     1,
				Name:   "abhinand",
				Email:  "abhi@gmail.com",
				Mobile: "9072001341",
			},
			expectedError: nil,
		},
		{
			name: "alredy exits",
			input: req.UserReq{
				Name:     "abhinand",
				Email:    "abhi@gmail.com",
				Mobile:   "9072001341",
				Password: "123456789",
			},
			buildStub: func(userRepo mockRepo.MockUserRepository) {
				userRepo.EXPECT().UserSignUp(gomock.Any(),
					EqCreateUserParams(req.UserReq{
						Name:     "abhinand",
						Email:    "abhi@gmail.com",
						Mobile:   "9072001341",
						Password: "123456789",
					},
						"123456789")).
					Times(1).
					Return(res.UserData{},
						errors.New("user alredy exits"))
			},
			expectedOutput: res.UserData{},
			expectedError:  errors.New("user alredy exits"),
		},
	}
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub(*userRepo)
			actualUser, err := userUseCase.UserSignUp(context.TODO(), tt.input)
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, actualUser, tt.expectedOutput)
		})
	}

}
