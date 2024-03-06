// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store/atlas (interfaces: PrivateEndpointLister)

// Package atlas is a generated GoMock package.
package atlas

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

// MockPrivateEndpointLister is a mock of PrivateEndpointLister interface.
type MockPrivateEndpointLister struct {
	ctrl     *gomock.Controller
	recorder *MockPrivateEndpointListerMockRecorder
}

// MockPrivateEndpointListerMockRecorder is the mock recorder for MockPrivateEndpointLister.
type MockPrivateEndpointListerMockRecorder struct {
	mock *MockPrivateEndpointLister
}

// NewMockPrivateEndpointLister creates a new mock instance.
func NewMockPrivateEndpointLister(ctrl *gomock.Controller) *MockPrivateEndpointLister {
	mock := &MockPrivateEndpointLister{ctrl: ctrl}
	mock.recorder = &MockPrivateEndpointListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrivateEndpointLister) EXPECT() *MockPrivateEndpointListerMockRecorder {
	return m.recorder
}

// PrivateEndpoints mocks base method.
func (m *MockPrivateEndpointLister) PrivateEndpoints(arg0, arg1 string) ([]admin.EndpointService, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrivateEndpoints", arg0, arg1)
	ret0, _ := ret[0].([]admin.EndpointService)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrivateEndpoints indicates an expected call of PrivateEndpoints.
func (mr *MockPrivateEndpointListerMockRecorder) PrivateEndpoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrivateEndpoints", reflect.TypeOf((*MockPrivateEndpointLister)(nil).PrivateEndpoints), arg0, arg1)
}
