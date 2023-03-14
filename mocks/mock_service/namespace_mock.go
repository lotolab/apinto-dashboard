// Code generated by MockGen. DO NOT EDIT.
// Source: service/namespace.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	model "github.com/eolinker/apinto-dashboard/model/namespace-model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockINamespaceService is a mock of INamespaceService interface.
type MockINamespaceService struct {
	ctrl     *gomock.Controller
	recorder *MockINamespaceServiceMockRecorder
}

// MockINamespaceServiceMockRecorder is the mock recorder for MockINamespaceService.
type MockINamespaceServiceMockRecorder struct {
	mock *MockINamespaceService
}

// NewMockINamespaceService creates a new mock instance.
func NewMockINamespaceService(ctrl *gomock.Controller) *MockINamespaceService {
	mock := &MockINamespaceService{ctrl: ctrl}
	mock.recorder = &MockINamespaceServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockINamespaceService) EXPECT() *MockINamespaceServiceMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockINamespaceService) GetAll() ([]*model.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*model.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockINamespaceServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockINamespaceService)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockINamespaceService) GetById(namespaceId int) (*model.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", namespaceId)
	ret0, _ := ret[0].(*model.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockINamespaceServiceMockRecorder) GetById(namespaceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockINamespaceService)(nil).GetById), namespaceId)
}

// GetByName mocks base method.
func (m *MockINamespaceService) GetByName(namespace string) (*model.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", namespace)
	ret0, _ := ret[0].(*model.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName.
func (mr *MockINamespaceServiceMockRecorder) GetByName(namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockINamespaceService)(nil).GetByName), namespace)
}
