// Code generated by MockGen. DO NOT EDIT.
// Source: user/user.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	context "context"
	reflect "reflect"

	access "github.com/eolinker/apinto-dashboard/access"
	user_model "github.com/eolinker/apinto-dashboard/modules/user/user-model"
	gomock "github.com/golang/mock/gomock"
)

// MockIUserInfoService is a mock of IUserInfoService interface.
type MockIUserInfoService struct {
	ctrl     *gomock.Controller
	recorder *MockIUserInfoServiceMockRecorder
}

// MockIUserInfoServiceMockRecorder is the mock recorder for MockIUserInfoService.
type MockIUserInfoServiceMockRecorder struct {
	mock *MockIUserInfoService
}

// NewMockIUserInfoService creates a new mock instance.
func NewMockIUserInfoService(ctrl *gomock.Controller) *MockIUserInfoService {
	mock := &MockIUserInfoService{ctrl: ctrl}
	mock.recorder = &MockIUserInfoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserInfoService) EXPECT() *MockIUserInfoServiceMockRecorder {
	return m.recorder
}

// GetAccessInfo mocks base method.
func (m *MockIUserInfoService) GetAccessInfo(ctx context.Context, userId int) (map[access.Access]struct{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessInfo", ctx, userId)
	ret0, _ := ret[0].(map[access.Access]struct{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessInfo indicates an expected call of GetAccessInfo.
func (mr *MockIUserInfoServiceMockRecorder) GetAccessInfo(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessInfo", reflect.TypeOf((*MockIUserInfoService)(nil).GetAccessInfo), ctx, userId)
}

// GetUserInfo mocks base method.
func (m *MockIUserInfoService) GetUserInfo(ctx context.Context, userId int) (*user_model.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfo", ctx, userId)
	ret0, _ := ret[0].(*user_model.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfo indicates an expected call of GetUserInfo.
func (mr *MockIUserInfoServiceMockRecorder) GetUserInfo(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfo", reflect.TypeOf((*MockIUserInfoService)(nil).GetUserInfo), ctx, userId)
}

// GetUserInfoMaps mocks base method.
func (m *MockIUserInfoService) GetUserInfoMaps(ctx context.Context, userId ...int) (map[int]*user_model.UserInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range userId {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserInfoMaps", varargs...)
	ret0, _ := ret[0].(map[int]*user_model.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfoMaps indicates an expected call of GetUserInfoMaps.
func (mr *MockIUserInfoServiceMockRecorder) GetUserInfoMaps(ctx interface{}, userId ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, userId...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfoMaps", reflect.TypeOf((*MockIUserInfoService)(nil).GetUserInfoMaps), varargs...)
}
