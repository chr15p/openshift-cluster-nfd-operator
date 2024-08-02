// Code generated by MockGen. DO NOT EDIT.
// Source: daemonset.go
//
// Generated by this command:
//
//	mockgen -source=daemonset.go -package=daemonset -destination=mock_daemonset.go DaemonsetAPI
//
// Package daemonset is a generated GoMock package.
package daemonset

import (
	context "context"
	reflect "reflect"

	v1 "github.com/openshift/cluster-nfd-operator/api/v1"
	gomock "go.uber.org/mock/gomock"
	v10 "k8s.io/api/apps/v1"
)

// MockDaemonsetAPI is a mock of DaemonsetAPI interface.
type MockDaemonsetAPI struct {
	ctrl     *gomock.Controller
	recorder *MockDaemonsetAPIMockRecorder
}

// MockDaemonsetAPIMockRecorder is the mock recorder for MockDaemonsetAPI.
type MockDaemonsetAPIMockRecorder struct {
	mock *MockDaemonsetAPI
}

// NewMockDaemonsetAPI creates a new mock instance.
func NewMockDaemonsetAPI(ctrl *gomock.Controller) *MockDaemonsetAPI {
	mock := &MockDaemonsetAPI{ctrl: ctrl}
	mock.recorder = &MockDaemonsetAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDaemonsetAPI) EXPECT() *MockDaemonsetAPIMockRecorder {
	return m.recorder
}

// DeleteDaemonSet mocks base method.
func (m *MockDaemonsetAPI) DeleteDaemonSet(ctx context.Context, namespace, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDaemonSet", ctx, namespace, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDaemonSet indicates an expected call of DeleteDaemonSet.
func (mr *MockDaemonsetAPIMockRecorder) DeleteDaemonSet(ctx, namespace, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDaemonSet", reflect.TypeOf((*MockDaemonsetAPI)(nil).DeleteDaemonSet), ctx, namespace, name)
}

// GetDaemonSet mocks base method.
func (m *MockDaemonsetAPI) GetDaemonSet(ctx context.Context, namespace, name string) (*v10.DaemonSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDaemonSet", ctx, namespace, name)
	ret0, _ := ret[0].(*v10.DaemonSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDaemonSet indicates an expected call of GetDaemonSet.
func (mr *MockDaemonsetAPIMockRecorder) GetDaemonSet(ctx, namespace, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDaemonSet", reflect.TypeOf((*MockDaemonsetAPI)(nil).GetDaemonSet), ctx, namespace, name)
}

// SetTopologyDaemonsetAsDesired mocks base method.
func (m *MockDaemonsetAPI) SetTopologyDaemonsetAsDesired(ctx context.Context, nfdInstance *v1.NodeFeatureDiscovery, topologyDS *v10.DaemonSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTopologyDaemonsetAsDesired", ctx, nfdInstance, topologyDS)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTopologyDaemonsetAsDesired indicates an expected call of SetTopologyDaemonsetAsDesired.
func (mr *MockDaemonsetAPIMockRecorder) SetTopologyDaemonsetAsDesired(ctx, nfdInstance, topologyDS any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTopologyDaemonsetAsDesired", reflect.TypeOf((*MockDaemonsetAPI)(nil).SetTopologyDaemonsetAsDesired), ctx, nfdInstance, topologyDS)
}

// SetWorkerDaemonsetAsDesired mocks base method.
func (m *MockDaemonsetAPI) SetWorkerDaemonsetAsDesired(ctx context.Context, nfdInstance *v1.NodeFeatureDiscovery, workerDS *v10.DaemonSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWorkerDaemonsetAsDesired", ctx, nfdInstance, workerDS)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWorkerDaemonsetAsDesired indicates an expected call of SetWorkerDaemonsetAsDesired.
func (mr *MockDaemonsetAPIMockRecorder) SetWorkerDaemonsetAsDesired(ctx, nfdInstance, workerDS any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWorkerDaemonsetAsDesired", reflect.TypeOf((*MockDaemonsetAPI)(nil).SetWorkerDaemonsetAsDesired), ctx, nfdInstance, workerDS)
}
