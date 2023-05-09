// Code generated by MockGen. DO NOT EDIT.
// Source: service/service.go

// Package mock_service is a generated GoMock package.
package mock

import (
	context "context"
	"github.com/eolinker/apinto-dashboard/modules/base/frontend-model"
	model "github.com/eolinker/apinto-dashboard/modules/strategy/strategy-model"
	upstream_model "github.com/eolinker/apinto-dashboard/modules/upstream/model"
	dto "github.com/eolinker/apinto-dashboard/modules/upstream/upstream-dto"
	upstream_entry2 "github.com/eolinker/apinto-dashboard/modules/upstream/upstream-entry"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIService is a mock of IService interface.
type MockIService struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceMockRecorder
}

// MockIServiceMockRecorder is the mock recorder for MockIService.
type MockIServiceMockRecorder struct {
	mock *MockIService
}

// NewMockIService creates a new mock instance.
func NewMockIService(ctrl *gomock.Controller) *MockIService {
	mock := &MockIService{ctrl: ctrl}
	mock.recorder = &MockIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIService) EXPECT() *MockIServiceMockRecorder {
	return m.recorder
}

// CreateService mocks base method.
func (m *MockIService) CreateService(ctx context.Context, namespaceID, userId int, input *dto.ServiceInfo, variableList []string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateService", ctx, namespaceID, userId, input, variableList)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateService indicates an expected call of CreateService.
func (mr *MockIServiceMockRecorder) CreateService(ctx, namespaceID, userId, input, variableList interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockIService)(nil).CreateService), ctx, namespaceID, userId, input, variableList)
}

// DeleteService mocks base method.
func (m *MockIService) DeleteService(ctx context.Context, namespaceID, userId int, serviceName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteService", ctx, namespaceID, userId, serviceName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteService indicates an expected call of DeleteService.
func (mr *MockIServiceMockRecorder) DeleteService(ctx, namespaceID, userId, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteService", reflect.TypeOf((*MockIService)(nil).DeleteService), ctx, namespaceID, userId, serviceName)
}

// GetLatestServiceVersion mocks base method.
func (m *MockIService) GetLatestServiceVersion(ctx context.Context, serviceID int) (*upstream_entry2.ServiceVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestServiceVersion", ctx, serviceID)
	ret0, _ := ret[0].(*upstream_entry2.ServiceVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestServiceVersion indicates an expected call of GetLatestServiceVersion.
func (mr *MockIServiceMockRecorder) GetLatestServiceVersion(ctx, serviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestServiceVersion", reflect.TypeOf((*MockIService)(nil).GetLatestServiceVersion), ctx, serviceID)
}

// GetServiceEnum mocks base method.
func (m *MockIService) GetServiceEnum(ctx context.Context, namespaceID int, searchName string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceEnum", ctx, namespaceID, searchName)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceEnum indicates an expected call of GetServiceEnum.
func (mr *MockIServiceMockRecorder) GetServiceEnum(ctx, namespaceID, searchName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceEnum", reflect.TypeOf((*MockIService)(nil).GetServiceEnum), ctx, namespaceID, searchName)
}

// GetServiceIDByName mocks base method.
func (m *MockIService) GetServiceIDByName(ctx context.Context, namespaceId int, serviceName string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceIDByName", ctx, namespaceId, serviceName)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceIDByName indicates an expected call of GetServiceIDByName.
func (mr *MockIServiceMockRecorder) GetServiceIDByName(ctx, namespaceId, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceIDByName", reflect.TypeOf((*MockIService)(nil).GetServiceIDByName), ctx, namespaceId, serviceName)
}

// GetServiceInfo mocks base method.
func (m *MockIService) GetServiceInfo(ctx context.Context, namespaceID int, serviceName string) (*upstream_model.ServiceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceInfo", ctx, namespaceID, serviceName)
	ret0, _ := ret[0].(*upstream_model.ServiceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceInfo indicates an expected call of GetServiceInfo.
func (mr *MockIServiceMockRecorder) GetServiceInfo(ctx, namespaceID, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceInfo", reflect.TypeOf((*MockIService)(nil).GetServiceInfo), ctx, namespaceID, serviceName)
}

// GetServiceList mocks base method.
func (m *MockIService) GetServiceList(ctx context.Context, namespaceID int, searchName string, pageNum, pageSize int) ([]*upstream_model.ServiceListItem, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceList", ctx, namespaceID, searchName, pageNum, pageSize)
	ret0, _ := ret[0].([]*upstream_model.ServiceListItem)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetServiceList indicates an expected call of GetServiceList.
func (mr *MockIServiceMockRecorder) GetServiceList(ctx, namespaceID, searchName, pageNum, pageSize interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceList", reflect.TypeOf((*MockIService)(nil).GetServiceList), ctx, namespaceID, searchName, pageNum, pageSize)
}

// GetServiceListAll mocks base method.
func (m *MockIService) GetServiceListAll(ctx context.Context, namespaceID int) ([]*upstream_model.ServiceListItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceListAll", ctx, namespaceID)
	ret0, _ := ret[0].([]*upstream_model.ServiceListItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceListAll indicates an expected call of GetServiceListAll.
func (mr *MockIServiceMockRecorder) GetServiceListAll(ctx, namespaceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceListAll", reflect.TypeOf((*MockIService)(nil).GetServiceListAll), ctx, namespaceID)
}

// GetServiceListByNames mocks base method.
func (m *MockIService) GetServiceListByNames(ctx context.Context, namespaceID int, names []string) ([]*upstream_model.ServiceListItem, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceListByNames", ctx, namespaceID, names)
	ret0, _ := ret[0].([]*upstream_model.ServiceListItem)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceListByNames indicates an expected call of GetServiceListByNames.
func (mr *MockIServiceMockRecorder) GetServiceListByNames(ctx, namespaceID, names interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceListByNames", reflect.TypeOf((*MockIService)(nil).GetServiceListByNames), ctx, namespaceID, names)
}

// GetServiceRemoteByNames mocks base method.
func (m *MockIService) GetServiceRemoteByNames(ctx context.Context, namespaceID int, uuids []string) ([]*model.RemoteServices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceRemoteByNames", ctx, namespaceID, uuids)
	ret0, _ := ret[0].([]*model.RemoteServices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceRemoteByNames indicates an expected call of GetServiceRemoteByNames.
func (mr *MockIServiceMockRecorder) GetServiceRemoteByNames(ctx, namespaceID, uuids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceRemoteByNames", reflect.TypeOf((*MockIService)(nil).GetServiceRemoteByNames), ctx, namespaceID, uuids)
}

// GetServiceRemoteOptions mocks base method.
func (m *MockIService) GetServiceRemoteOptions(ctx context.Context, namespaceID, pageNum, pageSize int, keyword string) ([]*model.RemoteServices, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceRemoteOptions", ctx, namespaceID, pageNum, pageSize, keyword)
	ret0, _ := ret[0].([]*model.RemoteServices)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetServiceRemoteOptions indicates an expected call of GetServiceRemoteOptions.
func (mr *MockIServiceMockRecorder) GetServiceRemoteOptions(ctx, namespaceID, pageNum, pageSize, keyword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceRemoteOptions", reflect.TypeOf((*MockIService)(nil).GetServiceRemoteOptions), ctx, namespaceID, pageNum, pageSize, keyword)
}

// GetServiceSchemaInfo mocks base method.
func (m *MockIService) GetServiceSchemaInfo(ctx context.Context, serviceID int) (*upstream_entry2.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceSchemaInfo", ctx, serviceID)
	ret0, _ := ret[0].(*upstream_entry2.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceSchemaInfo indicates an expected call of GetServiceSchemaInfo.
func (mr *MockIServiceMockRecorder) GetServiceSchemaInfo(ctx, serviceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceSchemaInfo", reflect.TypeOf((*MockIService)(nil).GetServiceSchemaInfo), ctx, serviceID)
}

// IsOnline mocks base method.
func (m *MockIService) IsOnline(ctx context.Context, clusterId, serviceId int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOnline", ctx, clusterId, serviceId)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsOnline indicates an expected call of IsOnline.
func (mr *MockIServiceMockRecorder) IsOnline(ctx, clusterId, serviceId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOnline", reflect.TypeOf((*MockIService)(nil).IsOnline), ctx, clusterId, serviceId)
}

// OfflineService mocks base method.
func (m *MockIService) OfflineService(ctx context.Context, namespaceId, operator int, serviceName, clusterName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OfflineService", ctx, namespaceId, operator, serviceName, clusterName)
	ret0, _ := ret[0].(error)
	return ret0
}

// OfflineService indicates an expected call of OfflineService.
func (mr *MockIServiceMockRecorder) OfflineService(ctx, namespaceId, operator, serviceName, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfflineService", reflect.TypeOf((*MockIService)(nil).OfflineService), ctx, namespaceId, operator, serviceName, clusterName)
}

// OnlineList mocks base method.
func (m *MockIService) OnlineList(ctx context.Context, namespaceId int, serviceName string) ([]*upstream_model.ServiceOnline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnlineList", ctx, namespaceId, serviceName)
	ret0, _ := ret[0].([]*upstream_model.ServiceOnline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OnlineList indicates an expected call of OnlineList.
func (mr *MockIServiceMockRecorder) OnlineList(ctx, namespaceId, serviceName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnlineList", reflect.TypeOf((*MockIService)(nil).OnlineList), ctx, namespaceId, serviceName)
}

// OnlineService mocks base method.
func (m *MockIService) OnlineService(ctx context.Context, namespaceId, operator int, serviceName, clusterName string) (*frontend_model.Router, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnlineService", ctx, namespaceId, operator, serviceName, clusterName)
	ret0, _ := ret[0].(*frontend_model.Router)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OnlineService indicates an expected call of OnlineService.
func (mr *MockIServiceMockRecorder) OnlineService(ctx, namespaceId, operator, serviceName, clusterName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnlineService", reflect.TypeOf((*MockIService)(nil).OnlineService), ctx, namespaceId, operator, serviceName, clusterName)
}

// UpdateService mocks base method.
func (m *MockIService) UpdateService(ctx context.Context, namespaceID, userId int, input *dto.ServiceInfo, variableList []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateService", ctx, namespaceID, userId, input, variableList)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateService indicates an expected call of UpdateService.
func (mr *MockIServiceMockRecorder) UpdateService(ctx, namespaceID, userId, input, variableList interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockIService)(nil).UpdateService), ctx, namespaceID, userId, input, variableList)
}
