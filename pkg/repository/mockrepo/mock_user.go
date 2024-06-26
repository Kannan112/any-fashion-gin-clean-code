// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interface/user.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	req "github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	res "github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	domain "github.com/kannan112/go-gin-clean-arch/pkg/domain"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// AccountVerify mocks base method.
func (m *MockUserRepository) AccountVerify(phone string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccountVerify", phone)
	ret0, _ := ret[0].(error)
	return ret0
}

// AccountVerify indicates an expected call of AccountVerify.
func (mr *MockUserRepositoryMockRecorder) AccountVerify(phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccountVerify", reflect.TypeOf((*MockUserRepository)(nil).AccountVerify), phone)
}

// AddAddress mocks base method.
func (m *MockUserRepository) AddAddress(id int, address req.AddAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", id, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockUserRepositoryMockRecorder) AddAddress(id, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockUserRepository)(nil).AddAddress), id, address)
}

// AuthLogin mocks base method.
func (m *MockUserRepository) AuthLogin(email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthLogin", email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthLogin indicates an expected call of AuthLogin.
func (mr *MockUserRepositoryMockRecorder) AuthLogin(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthLogin", reflect.TypeOf((*MockUserRepository)(nil).AuthLogin), email)
}

// AuthSignUp mocks base method.
func (m *MockUserRepository) AuthSignUp(Oauth req.GoogleAuth) (res.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthSignUp", Oauth)
	ret0, _ := ret[0].(res.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthSignUp indicates an expected call of AuthSignUp.
func (mr *MockUserRepositoryMockRecorder) AuthSignUp(Oauth interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthSignUp", reflect.TypeOf((*MockUserRepository)(nil).AuthSignUp), Oauth)
}

// CheckVerifyPhone mocks base method.
func (m *MockUserRepository) CheckVerifyPhone(mobileNo string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckVerifyPhone", mobileNo)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckVerifyPhone indicates an expected call of CheckVerifyPhone.
func (mr *MockUserRepositoryMockRecorder) CheckVerifyPhone(mobileNo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckVerifyPhone", reflect.TypeOf((*MockUserRepository)(nil).CheckVerifyPhone), mobileNo)
}

// DeleteAddress mocks base method.
func (m *MockUserRepository) DeleteAddress(ctx context.Context, userId, AddressesId int) ([]domain.Addresss, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddress", ctx, userId, AddressesId)
	ret0, _ := ret[0].([]domain.Addresss)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAddress indicates an expected call of DeleteAddress.
func (mr *MockUserRepositoryMockRecorder) DeleteAddress(ctx, userId, AddressesId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddress", reflect.TypeOf((*MockUserRepository)(nil).DeleteAddress), ctx, userId, AddressesId)
}

// EditProfile mocks base method.
func (m *MockUserRepository) EditProfile(id int, profile req.UserReq) (res.UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditProfile", id, profile)
	ret0, _ := ret[0].(res.UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditProfile indicates an expected call of EditProfile.
func (mr *MockUserRepositoryMockRecorder) EditProfile(id, profile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditProfile", reflect.TypeOf((*MockUserRepository)(nil).EditProfile), id, profile)
}

// FindAddress mocks base method.
func (m *MockUserRepository) FindAddress(ctx context.Context, userId int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAddress", ctx, userId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAddress indicates an expected call of FindAddress.
func (mr *MockUserRepositoryMockRecorder) FindAddress(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAddress", reflect.TypeOf((*MockUserRepository)(nil).FindAddress), ctx, userId)
}

// GetUserDetailsFromUserID mocks base method.
func (m *MockUserRepository) GetUserDetailsFromUserID(userId uint) (domain.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserDetailsFromUserID", userId)
	ret0, _ := ret[0].(domain.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserDetailsFromUserID indicates an expected call of GetUserDetailsFromUserID.
func (mr *MockUserRepositoryMockRecorder) GetUserDetailsFromUserID(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserDetailsFromUserID", reflect.TypeOf((*MockUserRepository)(nil).GetUserDetailsFromUserID), userId)
}

// IsSignIn mocks base method.
func (m *MockUserRepository) IsSignIn(phno string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSignIn", phno)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSignIn indicates an expected call of IsSignIn.
func (mr *MockUserRepositoryMockRecorder) IsSignIn(phno interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSignIn", reflect.TypeOf((*MockUserRepository)(nil).IsSignIn), phno)
}

// ListAllAddress mocks base method.
func (m *MockUserRepository) ListAllAddress(id int) ([]domain.Addresss, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllAddress", id)
	ret0, _ := ret[0].([]domain.Addresss)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllAddress indicates an expected call of ListAllAddress.
func (mr *MockUserRepositoryMockRecorder) ListAllAddress(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllAddress", reflect.TypeOf((*MockUserRepository)(nil).ListAllAddress), id)
}

// UpdateAddress mocks base method.
func (m *MockUserRepository) UpdateAddress(id, addressId int, address req.AddAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddress", id, addressId, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAddress indicates an expected call of UpdateAddress.
func (mr *MockUserRepositoryMockRecorder) UpdateAddress(id, addressId, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddress", reflect.TypeOf((*MockUserRepository)(nil).UpdateAddress), id, addressId, address)
}

// UserLogin mocks base method.
func (m *MockUserRepository) UserLogin(ctx context.Context, email string) (domain.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserLogin", ctx, email)
	ret0, _ := ret[0].(domain.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserLogin indicates an expected call of UserLogin.
func (mr *MockUserRepositoryMockRecorder) UserLogin(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserLogin", reflect.TypeOf((*MockUserRepository)(nil).UserLogin), ctx, email)
}

// UserSignUp mocks base method.
func (m *MockUserRepository) UserSignUp(ctx context.Context, user req.UserReq) (res.UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignUp", ctx, user)
	ret0, _ := ret[0].(res.UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignUp indicates an expected call of UserSignUp.
func (mr *MockUserRepositoryMockRecorder) UserSignUp(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignUp", reflect.TypeOf((*MockUserRepository)(nil).UserSignUp), ctx, user)
}

// ViewProfile mocks base method.
func (m *MockUserRepository) ViewProfile(id int) (res.UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewProfile", id)
	ret0, _ := ret[0].(res.UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewProfile indicates an expected call of ViewProfile.
func (mr *MockUserRepositoryMockRecorder) ViewProfile(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewProfile", reflect.TypeOf((*MockUserRepository)(nil).ViewProfile), id)
}
