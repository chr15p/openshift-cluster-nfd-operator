// Code generated by MockGen. DO NOT EDIT.
// Source: status.go
//
// Generated by this command:
//
//	mockgen -source=status.go -package=status -destination=mock_status.go statusHelperAPI
//
// Package status is a generated GoMock package.
package status

import (
	context "context"
	reflect "reflect"

	v1 "github.com/openshift/cluster-nfd-operator/api/v1"
	gomock "go.uber.org/mock/gomock"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MockStatusAPI is a mock of StatusAPI interface.
type MockStatusAPI struct {
	ctrl     *gomock.Controller
	recorder *MockStatusAPIMockRecorder
}

// MockStatusAPIMockRecorder is the mock recorder for MockStatusAPI.
type MockStatusAPIMockRecorder struct {
	mock *MockStatusAPI
}

// NewMockStatusAPI creates a new mock instance.
func NewMockStatusAPI(ctrl *gomock.Controller) *MockStatusAPI {
	mock := &MockStatusAPI{ctrl: ctrl}
	mock.recorder = &MockStatusAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatusAPI) EXPECT() *MockStatusAPIMockRecorder {
	return m.recorder
}

// AreConditionsEqual mocks base method.
func (m *MockStatusAPI) AreConditionsEqual(prevConditions, newConditions []v10.Condition) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AreConditionsEqual", prevConditions, newConditions)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AreConditionsEqual indicates an expected call of AreConditionsEqual.
func (mr *MockStatusAPIMockRecorder) AreConditionsEqual(prevConditions, newConditions any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AreConditionsEqual", reflect.TypeOf((*MockStatusAPI)(nil).AreConditionsEqual), prevConditions, newConditions)
}

// GetConditions mocks base method.
func (m *MockStatusAPI) GetConditions(ctx context.Context, nfdInstance *v1.NodeFeatureDiscovery) []v10.Condition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConditions", ctx, nfdInstance)
	ret0, _ := ret[0].([]v10.Condition)
	return ret0
}

// GetConditions indicates an expected call of GetConditions.
func (mr *MockStatusAPIMockRecorder) GetConditions(ctx, nfdInstance any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConditions", reflect.TypeOf((*MockStatusAPI)(nil).GetConditions), ctx, nfdInstance)
}

// MockstatusHelperAPI is a mock of statusHelperAPI interface.
type MockstatusHelperAPI struct {
	ctrl     *gomock.Controller
	recorder *MockstatusHelperAPIMockRecorder
}

// MockstatusHelperAPIMockRecorder is the mock recorder for MockstatusHelperAPI.
type MockstatusHelperAPIMockRecorder struct {
	mock *MockstatusHelperAPI
}

// NewMockstatusHelperAPI creates a new mock instance.
func NewMockstatusHelperAPI(ctrl *gomock.Controller) *MockstatusHelperAPI {
	mock := &MockstatusHelperAPI{ctrl: ctrl}
	mock.recorder = &MockstatusHelperAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockstatusHelperAPI) EXPECT() *MockstatusHelperAPIMockRecorder {
	return m.recorder
}

// getGCNotAvailableConditions mocks base method.
func (m *MockstatusHelperAPI) getGCNotAvailableConditions(ctx context.Context, nfdInstance *v1.NodeFeatureDiscovery) []v10.Condition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getGCNotAvailableConditions", ctx, nfdInstance)
	ret0, _ := ret[0].([]v10.Condition)
	return ret0
}

// getGCNotAvailableConditions indicates an expected call of getGCNotAvailableConditions.
func (mr *MockstatusHelperAPIMockRecorder) getGCNotAvailableConditions(ctx, nfdInstance any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getGCNotAvailableConditions", reflect.TypeOf((*MockstatusHelperAPI)(nil).getGCNotAvailableConditions), ctx, nfdInstance)
}

// getMasterNotAvailableConditions mocks base method.
func (m *MockstatusHelperAPI) getMasterNotAvailableConditions(ctx context.Context, nfdInstance *v1.NodeFeatureDiscovery) []v10.Condition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getMasterNotAvailableConditions", ctx, nfdInstance)
	ret0, _ := ret[0].([]v10.Condition)
	return ret0
}

// getMasterNotAvailableConditions indicates an expected call of getMasterNotAvailableConditions.
func (mr *MockstatusHelperAPIMockRecorder) getMasterNotAvailableConditions(ctx, nfdInstance any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getMasterNotAvailableConditions", reflect.TypeOf((*MockstatusHelperAPI)(nil).getMasterNotAvailableConditions), ctx, nfdInstance)
}

// getTopologyNotAvailableConditions mocks base method.
func (m *MockstatusHelperAPI) getTopologyNotAvailableConditions(ctx context.Context, nfdInstance *v1.NodeFeatureDiscovery) []v10.Condition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getTopologyNotAvailableConditions", ctx, nfdInstance)
	ret0, _ := ret[0].([]v10.Condition)
	return ret0
}

// getTopologyNotAvailableConditions indicates an expected call of getTopologyNotAvailableConditions.
func (mr *MockstatusHelperAPIMockRecorder) getTopologyNotAvailableConditions(ctx, nfdInstance any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getTopologyNotAvailableConditions", reflect.TypeOf((*MockstatusHelperAPI)(nil).getTopologyNotAvailableConditions), ctx, nfdInstance)
}

// getWorkerNotAvailableConditions mocks base method.
func (m *MockstatusHelperAPI) getWorkerNotAvailableConditions(ctx context.Context, nfdInstance *v1.NodeFeatureDiscovery) []v10.Condition {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getWorkerNotAvailableConditions", ctx, nfdInstance)
	ret0, _ := ret[0].([]v10.Condition)
	return ret0
}

// getWorkerNotAvailableConditions indicates an expected call of getWorkerNotAvailableConditions.
func (mr *MockstatusHelperAPIMockRecorder) getWorkerNotAvailableConditions(ctx, nfdInstance any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getWorkerNotAvailableConditions", reflect.TypeOf((*MockstatusHelperAPI)(nil).getWorkerNotAvailableConditions), ctx, nfdInstance)
}
