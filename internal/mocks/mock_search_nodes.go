// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: SearchNodesLister,SearchNodesCreator,SearchNodesUpdater,SearchNodesDeleter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20231115007/admin"
)

// MockSearchNodesLister is a mock of SearchNodesLister interface.
type MockSearchNodesLister struct {
	ctrl     *gomock.Controller
	recorder *MockSearchNodesListerMockRecorder
}

// MockSearchNodesListerMockRecorder is the mock recorder for MockSearchNodesLister.
type MockSearchNodesListerMockRecorder struct {
	mock *MockSearchNodesLister
}

// NewMockSearchNodesLister creates a new mock instance.
func NewMockSearchNodesLister(ctrl *gomock.Controller) *MockSearchNodesLister {
	mock := &MockSearchNodesLister{ctrl: ctrl}
	mock.recorder = &MockSearchNodesListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchNodesLister) EXPECT() *MockSearchNodesListerMockRecorder {
	return m.recorder
}

// SearchNodes mocks base method.
func (m *MockSearchNodesLister) SearchNodes(arg0, arg1 string) (*admin.ApiSearchDeploymentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchNodes", arg0, arg1)
	ret0, _ := ret[0].(*admin.ApiSearchDeploymentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchNodes indicates an expected call of SearchNodes.
func (mr *MockSearchNodesListerMockRecorder) SearchNodes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchNodes", reflect.TypeOf((*MockSearchNodesLister)(nil).SearchNodes), arg0, arg1)
}

// MockSearchNodesCreator is a mock of SearchNodesCreator interface.
type MockSearchNodesCreator struct {
	ctrl     *gomock.Controller
	recorder *MockSearchNodesCreatorMockRecorder
}

// MockSearchNodesCreatorMockRecorder is the mock recorder for MockSearchNodesCreator.
type MockSearchNodesCreatorMockRecorder struct {
	mock *MockSearchNodesCreator
}

// NewMockSearchNodesCreator creates a new mock instance.
func NewMockSearchNodesCreator(ctrl *gomock.Controller) *MockSearchNodesCreator {
	mock := &MockSearchNodesCreator{ctrl: ctrl}
	mock.recorder = &MockSearchNodesCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchNodesCreator) EXPECT() *MockSearchNodesCreatorMockRecorder {
	return m.recorder
}

// CreateSearchNodes mocks base method.
func (m *MockSearchNodesCreator) CreateSearchNodes(arg0, arg1 string, arg2 *admin.ApiSearchDeploymentRequest) (*admin.ApiSearchDeploymentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSearchNodes", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.ApiSearchDeploymentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSearchNodes indicates an expected call of CreateSearchNodes.
func (mr *MockSearchNodesCreatorMockRecorder) CreateSearchNodes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSearchNodes", reflect.TypeOf((*MockSearchNodesCreator)(nil).CreateSearchNodes), arg0, arg1, arg2)
}

// SearchNodes mocks base method.
func (m *MockSearchNodesCreator) SearchNodes(arg0, arg1 string) (*admin.ApiSearchDeploymentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchNodes", arg0, arg1)
	ret0, _ := ret[0].(*admin.ApiSearchDeploymentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchNodes indicates an expected call of SearchNodes.
func (mr *MockSearchNodesCreatorMockRecorder) SearchNodes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchNodes", reflect.TypeOf((*MockSearchNodesCreator)(nil).SearchNodes), arg0, arg1)
}

// MockSearchNodesUpdater is a mock of SearchNodesUpdater interface.
type MockSearchNodesUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockSearchNodesUpdaterMockRecorder
}

// MockSearchNodesUpdaterMockRecorder is the mock recorder for MockSearchNodesUpdater.
type MockSearchNodesUpdaterMockRecorder struct {
	mock *MockSearchNodesUpdater
}

// NewMockSearchNodesUpdater creates a new mock instance.
func NewMockSearchNodesUpdater(ctrl *gomock.Controller) *MockSearchNodesUpdater {
	mock := &MockSearchNodesUpdater{ctrl: ctrl}
	mock.recorder = &MockSearchNodesUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchNodesUpdater) EXPECT() *MockSearchNodesUpdaterMockRecorder {
	return m.recorder
}

// SearchNodes mocks base method.
func (m *MockSearchNodesUpdater) SearchNodes(arg0, arg1 string) (*admin.ApiSearchDeploymentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchNodes", arg0, arg1)
	ret0, _ := ret[0].(*admin.ApiSearchDeploymentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchNodes indicates an expected call of SearchNodes.
func (mr *MockSearchNodesUpdaterMockRecorder) SearchNodes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchNodes", reflect.TypeOf((*MockSearchNodesUpdater)(nil).SearchNodes), arg0, arg1)
}

// UpdateSearchNodes mocks base method.
func (m *MockSearchNodesUpdater) UpdateSearchNodes(arg0, arg1 string, arg2 *admin.ApiSearchDeploymentRequest) (*admin.ApiSearchDeploymentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSearchNodes", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.ApiSearchDeploymentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSearchNodes indicates an expected call of UpdateSearchNodes.
func (mr *MockSearchNodesUpdaterMockRecorder) UpdateSearchNodes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSearchNodes", reflect.TypeOf((*MockSearchNodesUpdater)(nil).UpdateSearchNodes), arg0, arg1, arg2)
}

// MockSearchNodesDeleter is a mock of SearchNodesDeleter interface.
type MockSearchNodesDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockSearchNodesDeleterMockRecorder
}

// MockSearchNodesDeleterMockRecorder is the mock recorder for MockSearchNodesDeleter.
type MockSearchNodesDeleterMockRecorder struct {
	mock *MockSearchNodesDeleter
}

// NewMockSearchNodesDeleter creates a new mock instance.
func NewMockSearchNodesDeleter(ctrl *gomock.Controller) *MockSearchNodesDeleter {
	mock := &MockSearchNodesDeleter{ctrl: ctrl}
	mock.recorder = &MockSearchNodesDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearchNodesDeleter) EXPECT() *MockSearchNodesDeleterMockRecorder {
	return m.recorder
}

// DeleteSearchNodes mocks base method.
func (m *MockSearchNodesDeleter) DeleteSearchNodes(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSearchNodes", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSearchNodes indicates an expected call of DeleteSearchNodes.
func (mr *MockSearchNodesDeleterMockRecorder) DeleteSearchNodes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSearchNodes", reflect.TypeOf((*MockSearchNodesDeleter)(nil).DeleteSearchNodes), arg0, arg1)
}

// SearchNodes mocks base method.
func (m *MockSearchNodesDeleter) SearchNodes(arg0, arg1 string) (*admin.ApiSearchDeploymentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchNodes", arg0, arg1)
	ret0, _ := ret[0].(*admin.ApiSearchDeploymentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchNodes indicates an expected call of SearchNodes.
func (mr *MockSearchNodesDeleterMockRecorder) SearchNodes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchNodes", reflect.TypeOf((*MockSearchNodesDeleter)(nil).SearchNodes), arg0, arg1)
}
