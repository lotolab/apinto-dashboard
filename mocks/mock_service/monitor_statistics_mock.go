// Code generated by MockGen. DO NOT EDIT.
// Source: service/monitor_statistics.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	"github.com/eolinker/apinto-dashboard/model/monitor-model"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockIMonitorStatistics is a mock of IMonitorStatistics interface.
type MockIMonitorStatistics struct {
	ctrl     *gomock.Controller
	recorder *MockIMonitorStatisticsMockRecorder
}

// MockIMonitorStatisticsMockRecorder is the mock recorder for MockIMonitorStatistics.
type MockIMonitorStatisticsMockRecorder struct {
	mock *MockIMonitorStatistics
}

// NewMockIMonitorStatistics creates a new mock instance.
func NewMockIMonitorStatistics(ctrl *gomock.Controller) *MockIMonitorStatistics {
	mock := &MockIMonitorStatistics{ctrl: ctrl}
	mock.recorder = &MockIMonitorStatisticsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMonitorStatistics) EXPECT() *MockIMonitorStatisticsMockRecorder {
	return m.recorder
}

// CircularMap mocks base method.
func (m *MockIMonitorStatistics) CircularMap(ctx context.Context, namespaceId int, partitionId string, start, end time.Time, wheres []monitor_model.MonWhereItem) (*monitor_model.CircularDate, *monitor_model.CircularDate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CircularMap", ctx, namespaceId, partitionId, start, end, wheres)
	ret0, _ := ret[0].(*monitor_model.CircularDate)
	ret1, _ := ret[1].(*monitor_model.CircularDate)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CircularMap indicates an expected call of CircularMap.
func (mr *MockIMonitorStatisticsMockRecorder) CircularMap(ctx, namespaceId, partitionId, start, end, wheres interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CircularMap", reflect.TypeOf((*MockIMonitorStatistics)(nil).CircularMap), ctx, namespaceId, partitionId, start, end, wheres)
}

// IPTrend mocks base method.
func (m *MockIMonitorStatistics) IPTrend(ctx context.Context, namespaceId int, partitionId string, start, end time.Time, wheres []monitor_model.MonWhereItem) (*monitor_model.MonCallCountInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IPTrend", ctx, namespaceId, partitionId, start, end, wheres)
	ret0, _ := ret[0].(*monitor_model.MonCallCountInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IPTrend indicates an expected call of IPTrend.
func (mr *MockIMonitorStatisticsMockRecorder) IPTrend(ctx, namespaceId, partitionId, start, end, wheres interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IPTrend", reflect.TypeOf((*MockIMonitorStatistics)(nil).IPTrend), ctx, namespaceId, partitionId, start, end, wheres)
}

// MessageTrend mocks base method.
func (m *MockIMonitorStatistics) MessageTrend(ctx context.Context, namespaceId int, partitionId string, start, end time.Time, wheres []monitor_model.MonWhereItem) (*monitor_model.MessageTrend, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MessageTrend", ctx, namespaceId, partitionId, start, end, wheres)
	ret0, _ := ret[0].(*monitor_model.MessageTrend)
	ret1, _ := ret[1].(error)
	return ret0, "", ret1
}

// MessageTrend indicates an expected call of MessageTrend.
func (mr *MockIMonitorStatisticsMockRecorder) MessageTrend(ctx, namespaceId, partitionId, start, end, wheres interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MessageTrend", reflect.TypeOf((*MockIMonitorStatistics)(nil).MessageTrend), ctx, namespaceId, partitionId, start, end, wheres)
}

// ProxyStatistics mocks base method.
func (m *MockIMonitorStatistics) ProxyStatistics(ctx context.Context, namespaceId int, partitionId string, start, end time.Time, groupBy string, wheres []monitor_model.MonWhereItem, limit int) (map[string]monitor_model.MonCommonData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProxyStatistics", ctx, namespaceId, partitionId, start, end, groupBy, wheres, limit)
	ret0, _ := ret[0].(map[string]monitor_model.MonCommonData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProxyStatistics indicates an expected call of ProxyStatistics.
func (mr *MockIMonitorStatisticsMockRecorder) ProxyStatistics(ctx, namespaceId, partitionId, start, end, groupBy, wheres, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProxyStatistics", reflect.TypeOf((*MockIMonitorStatistics)(nil).ProxyStatistics), ctx, namespaceId, partitionId, start, end, groupBy, wheres, limit)
}

// ProxyTrend mocks base method.
func (m *MockIMonitorStatistics) ProxyTrend(ctx context.Context, namespaceId int, partitionId string, start, end time.Time, wheres []monitor_model.MonWhereItem) (*monitor_model.MonCallCountInfo, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProxyTrend", ctx, namespaceId, partitionId, start, end, wheres)
	ret0, _ := ret[0].(*monitor_model.MonCallCountInfo)
	ret1, _ := ret[1].(error)
	return ret0, "", ret1
}

// ProxyTrend indicates an expected call of ProxyTrend.
func (mr *MockIMonitorStatisticsMockRecorder) ProxyTrend(ctx, namespaceId, partitionId, start, end, wheres interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProxyTrend", reflect.TypeOf((*MockIMonitorStatistics)(nil).ProxyTrend), ctx, namespaceId, partitionId, start, end, wheres)
}

// Statistics mocks base method.
func (m *MockIMonitorStatistics) Statistics(ctx context.Context, namespaceId int, partitionId string, start, end time.Time, groupBy string, wheres []monitor_model.MonWhereItem, limit int) (map[string]monitor_model.MonCommonData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Statistics", ctx, namespaceId, partitionId, start, end, groupBy, wheres, limit)
	ret0, _ := ret[0].(map[string]monitor_model.MonCommonData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Statistics indicates an expected call of Statistics.
func (mr *MockIMonitorStatisticsMockRecorder) Statistics(ctx, namespaceId, partitionId, start, end, groupBy, wheres, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Statistics", reflect.TypeOf((*MockIMonitorStatistics)(nil).Statistics), ctx, namespaceId, partitionId, start, end, groupBy, wheres, limit)
}

// Trend mocks base method.
func (m *MockIMonitorStatistics) Trend(ctx context.Context, namespaceId int, partitionId string, start, end time.Time, wheres []monitor_model.MonWhereItem) (*monitor_model.MonCallCountInfo, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trend", ctx, namespaceId, partitionId, start, end, wheres)
	ret0, _ := ret[0].(*monitor_model.MonCallCountInfo)
	ret1, _ := ret[1].(error)
	return ret0, "", ret1
}

// Trend indicates an expected call of Trend.
func (mr *MockIMonitorStatisticsMockRecorder) Trend(ctx, namespaceId, partitionId, start, end, wheres interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trend", reflect.TypeOf((*MockIMonitorStatistics)(nil).Trend), ctx, namespaceId, partitionId, start, end, wheres)
}

// WarnStatistics mocks base method.
func (m *MockIMonitorStatistics) WarnStatistics(ctx context.Context, namespaceId int, partitionId string, start, end time.Time, group string, quotaType monitor_model.QuotaType, wheres []monitor_model.MonWhereItem) (map[string]float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WarnStatistics", ctx, namespaceId, partitionId, start, end, group, quotaType, wheres)
	ret0, _ := ret[0].(map[string]float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WarnStatistics indicates an expected call of WarnStatistics.
func (mr *MockIMonitorStatisticsMockRecorder) WarnStatistics(ctx, namespaceId, partitionId, start, end, group, quotaType, wheres interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WarnStatistics", reflect.TypeOf((*MockIMonitorStatistics)(nil).WarnStatistics), ctx, namespaceId, partitionId, start, end, group, quotaType, wheres)
}
