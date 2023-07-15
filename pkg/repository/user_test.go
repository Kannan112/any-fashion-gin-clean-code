package repository

import (
	"context"

	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"

	"github.com/stretchr/testify/assert"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserSignup(t *testing.T) {

	tests := []struct {
		name           string
		input          req.UserReq
		expectedOutput res.UserData
		buildStub      func(mock sqlmock.Sqlmock)
		expectedErr    error
	}{
		{
			name: "successful create user",

			input: req.UserReq{
				Name:     "alan",
				Email:    "alan@gmail.com",
				Mobile:   "+9195623999",
				Password: "alan@123",
			},
			expectedOutput: res.UserData{
				Id:     6,
				Name:   "alan",
				Email:  "alan@gmail.com",
				Mobile: "+9195623999",
			},

			buildStub: func(mock sqlmock.Sqlmock) {

				rows := sqlmock.NewRows([]string{"id", "name", "email", "mobile"}).
					AddRow(6, "alan", "alan@gmail.com", "+9195623999")

				mock.ExpectQuery("^INSERT INTO users (.+)$").
					WithArgs("alan", "alan@gmail.com", "+9195623999", "alan@123").
					WillReturnRows(rows)
			},
			expectedErr: nil,
		},
		{
			name: "duplicate user",
			input: req.UserReq{
				Name:     "alan",
				Email:    "alan@gmail.com",
				Mobile:   "+9195623999",
				Password: "alan@123",
			},
			expectedOutput: res.UserData{},
			buildStub: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^INSERT INTO users (.+)$").
					WithArgs("alan", "alan@gmail.com", "+9195623999", "alan@123").
					WillReturnError(errors.New("email or phone number alredy used"))
			},
			expectedErr: errors.New("email or phone number alredy used"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//New() method from sqlmock package create sqlmock database connection and a mock to manage expectations.
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			//initialize the db instance with the mock db connection
			gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
			if err != nil {
				t.Fatalf("an error '%s' was not expected when initializing a mock db session", err)
			}

			//create NewUserRepository mock by passing a pointer to gorm.DB
			userRepository := NewUserRepository(gormDB)

			// before we actually execute our function, we need to expect required DB actions
			tt.buildStub(mock)

			//call the actual method
			actualOutput, actualErr := userRepository.UserSignUp(context.TODO(), tt.input)
			// validate err is nil if we are not expecting to receive an error
			if tt.expectedErr == nil {
				assert.NoError(t, actualErr)
			} else { //validate whether expected and actual errors are same
				assert.Equal(t, tt.expectedErr, actualErr)
			}

			if !reflect.DeepEqual(tt.expectedOutput, actualOutput) {
				t.Errorf("got %v, but want %v", actualOutput, tt.expectedOutput)
			}

			// Check that all expectations were met
			err = mock.ExpectationsWereMet()
			if err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}
		})
	}

}

// func TestUserLogin(t *testing.T) {
// 	Logintest := []struct {
// 		name           string
// 		input          string
// 		expectedOutput domain.Users
// 		buildStub      func(mock sqlmock.Sqlmock)
// 		expectedError  error
// 	}{
// 		{
// 			name:  "login user",
// 			input: "abhinand@gmail.com",
// 			expectedOutput: domain.Users{
// 				ID:        1,
// 				Name:      "abhinand",
// 				Email:     "abhinand@gmail.com",
// 				Mobile:    "+919846789099",
// 				Password:  "HVAKHVIAKJEB",
// 				IsBlocked: false,
// 				CreatedAt: time.Time{},
// 			},
// 			buildStub: func(mock sqlmock.Sqlmock) {
// 				mock.ExpectQuery(`^SELECT EXISTS\(SELECT \* FROM users WHERE email=\$1 AND is_blocked=true\)$`).
// 					WithArgs("abhinand@gmail.com").
// 					WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

// 				mock.ExpectQuery(`SELECT \* FROM users WHERE email\=\?`).
// 					WithArgs("abhinand@gmail.com").
// 					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "mobile", "password", "is_blocked", "created_at"}).
// 						AddRow(1, "abhinand", "abhinand@gmail.com", "+919846789099", "HVAKHVIAKJEB", false, time.Time{}))
// 			},
// 			expectedError: nil,
// 		},
// 	}

// 	for _, tt := range Logintest {
// 		t.Run(tt.name, func(t *testing.T) {
// 			db, mock, err := sqlmock.New()
// 			if err != nil {
// 				t.Fatalf("an error %s was not expected while opening the database connection", err)
// 			}
// 			defer db.Close()

// 			gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
// 			if err != nil {
// 				t.Fatalf("an error %s was not expected while initializing the DB connection", err)
// 			}

// 			userRepository := NewUserRepository(gormDB)
// 			tt.buildStub(mock)
// 			actualOutput, actualErr := userRepository.UserLogin(context.TODO(), tt.input)

// 			if tt.expectedError == nil {
// 				assert.NoError(t, actualErr)
// 			} else {
// 				assert.EqualError(t, tt.expectedError, actualErr.Error())
// 			}

// 			if !reflect.DeepEqual(tt.expectedOutput, actualOutput) {
// 				t.Errorf("got %v, but want %v", actualOutput, tt.expectedOutput)
// 			}

// 			err = mock.ExpectationsWereMet()
// 			if err != nil {
// 				t.Errorf("unfulfilled expectation: %s", err)
// 			}
// 		})
// 	}
// }
