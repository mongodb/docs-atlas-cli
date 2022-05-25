// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/cli/auth (interfaces: Revoker,ConfigDeleter)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
)

// MockRevoker is a mock of Revoker interface.
type MockRevoker struct {
	ctrl     *gomock.Controller
	recorder *MockRevokerMockRecorder
}

// MockRevokerMockRecorder is the mock recorder for MockRevoker.
type MockRevokerMockRecorder struct {
	mock *MockRevoker
}

// NewMockRevoker creates a new mock instance.
func NewMockRevoker(ctrl *gomock.Controller) *MockRevoker {
	mock := &MockRevoker{ctrl: ctrl}
	mock.recorder = &MockRevokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRevoker) EXPECT() *MockRevokerMockRecorder {
	return m.recorder
}

// RevokeToken mocks base method.
func (m *MockRevoker) RevokeToken(arg0 context.Context, arg1, arg2 string) (*mongodbatlas.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeToken", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlas.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeToken indicates an expected call of RevokeToken.
func (mr *MockRevokerMockRecorder) RevokeToken(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeToken", reflect.TypeOf((*MockRevoker)(nil).RevokeToken), arg0, arg1, arg2)
}

// MockConfigDeleter is a mock of ConfigDeleter interface.
type MockConfigDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockConfigDeleterMockRecorder
}

// MockConfigDeleterMockRecorder is the mock recorder for MockConfigDeleter.
type MockConfigDeleterMockRecorder struct {
	mock *MockConfigDeleter
}

// NewMockConfigDeleter creates a new mock instance.
func NewMockConfigDeleter(ctrl *gomock.Controller) *MockConfigDeleter {
	mock := &MockConfigDeleter{ctrl: ctrl}
	mock.recorder = &MockConfigDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigDeleter) EXPECT() *MockConfigDeleterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockConfigDeleter) Delete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockConfigDeleterMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockConfigDeleter)(nil).Delete))
}
