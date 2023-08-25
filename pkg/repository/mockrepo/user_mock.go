// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interface/user.go

// Package mockrepo is a generated GoMock package.
package mockrepo

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/nadeem1815/premium-watch/pkg/domain"
	model "github.com/nadeem1815/premium-watch/pkg/utils/model"
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

// AddAddress mocks base method.
func (m *MockUserRepository) AddAddress(ctx context.Context, body model.AddressInput, userID string) (domain.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAddress", ctx, body, userID)
	ret0, _ := ret[0].(domain.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddAddress indicates an expected call of AddAddress.
func (mr *MockUserRepositoryMockRecorder) AddAddress(ctx, body, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAddress", reflect.TypeOf((*MockUserRepository)(nil).AddAddress), ctx, body, userID)
}

// BlockUser mocks base method.
func (m *MockUserRepository) BlockUser(ctx context.Context, user_id string) (domain.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockUser", ctx, user_id)
	ret0, _ := ret[0].(domain.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockUser indicates an expected call of BlockUser.
func (mr *MockUserRepositoryMockRecorder) BlockUser(ctx, user_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockUser", reflect.TypeOf((*MockUserRepository)(nil).BlockUser), ctx, user_id)
}

// FindByEmail mocks base method.
func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (model.UserLoginVarifier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", ctx, email)
	ret0, _ := ret[0].(model.UserLoginVarifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockUserRepositoryMockRecorder) FindByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserRepository)(nil).FindByEmail), ctx, email)
}

// UnBlockUser mocks base method.
func (m *MockUserRepository) UnBlockUser(ctx context.Context, user_id string) (domain.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnBlockUser", ctx, user_id)
	ret0, _ := ret[0].(domain.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnBlockUser indicates an expected call of UnBlockUser.
func (mr *MockUserRepositoryMockRecorder) UnBlockUser(ctx, user_id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnBlockUser", reflect.TypeOf((*MockUserRepository)(nil).UnBlockUser), ctx, user_id)
}

// UserRegister mocks base method.
func (m *MockUserRepository) UserRegister(ctx context.Context, user model.UsarDataInput) (model.UserDataOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserRegister", ctx, user)
	ret0, _ := ret[0].(model.UserDataOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UserRegister indicates an expected call of UserRegister.
func (mr *MockUserRepositoryMockRecorder) UserRegister(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserRegister", reflect.TypeOf((*MockUserRepository)(nil).UserRegister), ctx, user)
}

// ViewAddress mocks base method.
func (m *MockUserRepository) ViewAddress(ctx context.Context, userID string) (domain.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ViewAddress", ctx, userID)
	ret0, _ := ret[0].(domain.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ViewAddress indicates an expected call of ViewAddress.
func (mr *MockUserRepositoryMockRecorder) ViewAddress(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ViewAddress", reflect.TypeOf((*MockUserRepository)(nil).ViewAddress), ctx, userID)
}