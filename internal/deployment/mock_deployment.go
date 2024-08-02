// Code generated by MockGen. DO NOT EDIT.
// Source: deployment.go
//
// Generated by this command:
//
//	mockgen -source=deployment.go -package=deployment -destination=mock_deployment.go DeploymentAPI
//
// Package deployment is a generated GoMock package.
package deployment

import (
	context "context"
	reflect "reflect"

	v1 "github.com/openshift/cluster-nfd-operator/api/v1"
	gomock "go.uber.org/mock/gomock"
	v10 "k8s.io/api/apps/v1"
)

// MockDeploymentAPI is a mock of DeploymentAPI interface.
type MockDeploymentAPI struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentAPIMockRecorder
}

// MockDeploymentAPIMockRecorder is the mock recorder for MockDeploymentAPI.
type MockDeploymentAPIMockRecorder struct {
	mock *MockDeploymentAPI
}

// NewMockDeploymentAPI creates a new mock instance.
func NewMockDeploymentAPI(ctrl *gomock.Controller) *MockDeploymentAPI {
	mock := &MockDeploymentAPI{ctrl: ctrl}
	mock.recorder = &MockDeploymentAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentAPI) EXPECT() *MockDeploymentAPIMockRecorder {
	return m.recorder
}

// DeleteDeployment mocks base method.
func (m *MockDeploymentAPI) DeleteDeployment(ctx context.Context, namespace, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDeployment", ctx, namespace, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDeployment indicates an expected call of DeleteDeployment.
func (mr *MockDeploymentAPIMockRecorder) DeleteDeployment(ctx, namespace, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeployment", reflect.TypeOf((*MockDeploymentAPI)(nil).DeleteDeployment), ctx, namespace, name)
}

// GetDeployment mocks base method.
func (m *MockDeploymentAPI) GetDeployment(ctx context.Context, namespace, name string) (*v10.Deployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployment", ctx, namespace, name)
	ret0, _ := ret[0].(*v10.Deployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployment indicates an expected call of GetDeployment.
func (mr *MockDeploymentAPIMockRecorder) GetDeployment(ctx, namespace, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployment", reflect.TypeOf((*MockDeploymentAPI)(nil).GetDeployment), ctx, namespace, name)
}

// SetGCDeploymentAsDesired mocks base method.
func (m *MockDeploymentAPI) SetGCDeploymentAsDesired(nfdInstance *v1.NodeFeatureDiscovery, gcDep *v10.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetGCDeploymentAsDesired", nfdInstance, gcDep)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetGCDeploymentAsDesired indicates an expected call of SetGCDeploymentAsDesired.
func (mr *MockDeploymentAPIMockRecorder) SetGCDeploymentAsDesired(nfdInstance, gcDep any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGCDeploymentAsDesired", reflect.TypeOf((*MockDeploymentAPI)(nil).SetGCDeploymentAsDesired), nfdInstance, gcDep)
}

// SetMasterDeploymentAsDesired mocks base method.
func (m *MockDeploymentAPI) SetMasterDeploymentAsDesired(nfdInstance *v1.NodeFeatureDiscovery, masterDep *v10.Deployment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMasterDeploymentAsDesired", nfdInstance, masterDep)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMasterDeploymentAsDesired indicates an expected call of SetMasterDeploymentAsDesired.
func (mr *MockDeploymentAPIMockRecorder) SetMasterDeploymentAsDesired(nfdInstance, masterDep any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMasterDeploymentAsDesired", reflect.TypeOf((*MockDeploymentAPI)(nil).SetMasterDeploymentAsDesired), nfdInstance, masterDep)
}
