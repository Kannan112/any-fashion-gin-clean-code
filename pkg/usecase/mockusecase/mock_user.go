// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/usecase/interface/user.go

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

// MockUserUseCase is a mock of UserUseCase interface.
type MockUserUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUserUseCaseMockRecorder
}

// MockUserUseCaseMockRecorder is the mock recorder for MockUserUseCase.
type MockUserUseCaseMockRecorder struct {
	mock *MockUserUseCase
}

// NewMockUserUseCase creates a new mock instance.
func NewMockUserUseCase(ctrl *gomock.Controller) *MockUserUseCase {
	mock := &MockUserUseCase{ctrl: ctrl}
	mock.recorder = &MockUserUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUseCase) EXPECT() *MockUserUseCaseMockRecorder {
	return m.recorder
}

// AddAddress mocks base method.
func (m *MockUserUseCase) AddAddress(id int, body req.AddAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", id, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockUserUseCaseMockRecorder) AddAddress(id, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockUserUseCase)(nil).AddAddress), id, body)
}

// DeleteAddress mocks base method.
func (m *MockUserUseCase) DeleteAddress(ctx context.Context, userId, AddressesId int) ([]domain.Addresss, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAddress", ctx, userId, AddressesId)
	ret0, _ := ret[0].([]domain.Addresss)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAddress indicates an expected call of DeleteAddress.
func (mr *MockUserUseCaseMockRecorder) DeleteAddress(ctx, userId, AddressesId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAddress", reflect.TypeOf((*MockUserUseCase)(nil).DeleteAddress), ctx, userId, AddressesId)
}

// EditProfile mocks base method.
func (m *MockUserUseCase) EditProfile(id int, UpdateProfile req.UserReq) (res.UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditProfile", id, UpdateProfile)
	ret0, _ := ret[0].(res.UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditProfile indicates an expected call of EditProfile.
func (mr *MockUserUseCaseMockRecorder) EditProfile(id, UpdateProfile interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditProfile", reflect.TypeOf((*MockUserUseCase)(nil).EditProfile), id, UpdateProfile)
}

// FindAddress mocks base method.
func (m *MockUserUseCase) FindAddress(ctx context.Context, userId int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAddress", ctx, userId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAddress indicates an expected call of FindAddress.
func (mr *MockUserUseCaseMockRecorder) FindAddress(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAddress", reflect.TypeOf((*MockUserUseCase)(nil).FindAddress), ctx, userId)
}

// IsSignIn mocks base method.
func (m *MockUserUseCase) IsSignIn(phno string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSignIn", phno)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSignIn indicates an expected call of IsSignIn.
func (mr *MockUserUseCaseMockRecorder) IsSignIn(phno interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSignIn", reflect.TypeOf((*MockUserUseCase)(nil).IsSignIn), phno)
}

// ListallAddress mocks base method.
func (m *MockUserUseCase) ListallAddress(id int) ([]domain.Addresss, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListallAddress", id)
	ret0, _ := ret[0].([]domain.Addresss)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListallAddress indicates an expected call of ListallAddress.
func (mr *MockUserUseCaseMockRecorder) ListallAddress(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListallAddress", reflect.TypeOf((*MockUserUseCase)(nil).ListallAddress), id)
}

// UpdateAddress mocks base method.
func (m *MockUserUseCase) UpdateAddress(id, addressId int, address req.AddAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAddress", id, addressId, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAddress indicates an expected call of UpdateAddress.
func (mr *MockUserUseCaseMockRecorder) UpdateAddress(id, addressId, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAddress", reflect.TypeOf((*MockUserUseCase)(nil).UpdateAddress), id, addressId, address)
}

// UserLogin mocks base method.
func (m *MockUserUseCase) UserLogin(ctx context.Context, user req.LoginReq) (res.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserLogin", ctx, user)
	ret0, _ := ret[0].(res.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserLogin indicates an expected call of UserLogin.
func (mr *MockUserUseCaseMockRecorder) UserLogin(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserLogin", reflect.TypeOf((*MockUserUseCase)(nil).UserLogin), ctx, user)
}

// UserSignUp mocks base method.
func (m *MockUserUseCase) UserSignUp(ctx context.Context, user req.UserReq) (res.UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserSignUp", ctx, user)
	ret0, _ := ret[0].(res.UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserSignUp indicates an expected call of UserSignUp.
func (mr *MockUserUseCaseMockRecorder) UserSignUp(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserSignUp", reflect.TypeOf((*MockUserUseCase)(nil).UserSignUp), ctx, user)
}

// ViewProfile mocks base method.
func (m *MockUserUseCase) ViewProfile(id int) (res.UserData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewProfile", id)
	ret0, _ := ret[0].(res.UserData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewProfile indicates an expected call of ViewProfile.
func (mr *MockUserUseCaseMockRecorder) ViewProfile(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewProfile", reflect.TypeOf((*MockUserUseCase)(nil).ViewProfile), id)
}
