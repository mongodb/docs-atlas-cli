// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store (interfaces: DatabaseRoleLister,DatabaseRoleCreator,DatabaseRoleDeleter,DatabaseRoleUpdater,DatabaseRoleDescriber)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	admin "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

// MockDatabaseRoleLister is a mock of DatabaseRoleLister interface.
type MockDatabaseRoleLister struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseRoleListerMockRecorder
}

// MockDatabaseRoleListerMockRecorder is the mock recorder for MockDatabaseRoleLister.
type MockDatabaseRoleListerMockRecorder struct {
	mock *MockDatabaseRoleLister
}

// NewMockDatabaseRoleLister creates a new mock instance.
func NewMockDatabaseRoleLister(ctrl *gomock.Controller) *MockDatabaseRoleLister {
	mock := &MockDatabaseRoleLister{ctrl: ctrl}
	mock.recorder = &MockDatabaseRoleListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseRoleLister) EXPECT() *MockDatabaseRoleListerMockRecorder {
	return m.recorder
}

// DatabaseRoles mocks base method.
func (m *MockDatabaseRoleLister) DatabaseRoles(arg0 string) ([]admin.UserCustomDBRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseRoles", arg0)
	ret0, _ := ret[0].([]admin.UserCustomDBRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatabaseRoles indicates an expected call of DatabaseRoles.
func (mr *MockDatabaseRoleListerMockRecorder) DatabaseRoles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseRoles", reflect.TypeOf((*MockDatabaseRoleLister)(nil).DatabaseRoles), arg0)
}

// MockDatabaseRoleCreator is a mock of DatabaseRoleCreator interface.
type MockDatabaseRoleCreator struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseRoleCreatorMockRecorder
}

// MockDatabaseRoleCreatorMockRecorder is the mock recorder for MockDatabaseRoleCreator.
type MockDatabaseRoleCreatorMockRecorder struct {
	mock *MockDatabaseRoleCreator
}

// NewMockDatabaseRoleCreator creates a new mock instance.
func NewMockDatabaseRoleCreator(ctrl *gomock.Controller) *MockDatabaseRoleCreator {
	mock := &MockDatabaseRoleCreator{ctrl: ctrl}
	mock.recorder = &MockDatabaseRoleCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseRoleCreator) EXPECT() *MockDatabaseRoleCreatorMockRecorder {
	return m.recorder
}

// CreateDatabaseRole mocks base method.
func (m *MockDatabaseRoleCreator) CreateDatabaseRole(arg0 string, arg1 *admin.UserCustomDBRole) (*admin.UserCustomDBRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDatabaseRole", arg0, arg1)
	ret0, _ := ret[0].(*admin.UserCustomDBRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDatabaseRole indicates an expected call of CreateDatabaseRole.
func (mr *MockDatabaseRoleCreatorMockRecorder) CreateDatabaseRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDatabaseRole", reflect.TypeOf((*MockDatabaseRoleCreator)(nil).CreateDatabaseRole), arg0, arg1)
}

// MockDatabaseRoleDeleter is a mock of DatabaseRoleDeleter interface.
type MockDatabaseRoleDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseRoleDeleterMockRecorder
}

// MockDatabaseRoleDeleterMockRecorder is the mock recorder for MockDatabaseRoleDeleter.
type MockDatabaseRoleDeleterMockRecorder struct {
	mock *MockDatabaseRoleDeleter
}

// NewMockDatabaseRoleDeleter creates a new mock instance.
func NewMockDatabaseRoleDeleter(ctrl *gomock.Controller) *MockDatabaseRoleDeleter {
	mock := &MockDatabaseRoleDeleter{ctrl: ctrl}
	mock.recorder = &MockDatabaseRoleDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseRoleDeleter) EXPECT() *MockDatabaseRoleDeleterMockRecorder {
	return m.recorder
}

// DeleteDatabaseRole mocks base method.
func (m *MockDatabaseRoleDeleter) DeleteDatabaseRole(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDatabaseRole", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDatabaseRole indicates an expected call of DeleteDatabaseRole.
func (mr *MockDatabaseRoleDeleterMockRecorder) DeleteDatabaseRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDatabaseRole", reflect.TypeOf((*MockDatabaseRoleDeleter)(nil).DeleteDatabaseRole), arg0, arg1)
}

// MockDatabaseRoleUpdater is a mock of DatabaseRoleUpdater interface.
type MockDatabaseRoleUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseRoleUpdaterMockRecorder
}

// MockDatabaseRoleUpdaterMockRecorder is the mock recorder for MockDatabaseRoleUpdater.
type MockDatabaseRoleUpdaterMockRecorder struct {
	mock *MockDatabaseRoleUpdater
}

// NewMockDatabaseRoleUpdater creates a new mock instance.
func NewMockDatabaseRoleUpdater(ctrl *gomock.Controller) *MockDatabaseRoleUpdater {
	mock := &MockDatabaseRoleUpdater{ctrl: ctrl}
	mock.recorder = &MockDatabaseRoleUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseRoleUpdater) EXPECT() *MockDatabaseRoleUpdaterMockRecorder {
	return m.recorder
}

// DatabaseRole mocks base method.
func (m *MockDatabaseRoleUpdater) DatabaseRole(arg0, arg1 string) (*admin.UserCustomDBRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseRole", arg0, arg1)
	ret0, _ := ret[0].(*admin.UserCustomDBRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatabaseRole indicates an expected call of DatabaseRole.
func (mr *MockDatabaseRoleUpdaterMockRecorder) DatabaseRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseRole", reflect.TypeOf((*MockDatabaseRoleUpdater)(nil).DatabaseRole), arg0, arg1)
}

// UpdateDatabaseRole mocks base method.
func (m *MockDatabaseRoleUpdater) UpdateDatabaseRole(arg0, arg1 string, arg2 *admin.UserCustomDBRole) (*admin.UserCustomDBRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDatabaseRole", arg0, arg1, arg2)
	ret0, _ := ret[0].(*admin.UserCustomDBRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDatabaseRole indicates an expected call of UpdateDatabaseRole.
func (mr *MockDatabaseRoleUpdaterMockRecorder) UpdateDatabaseRole(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDatabaseRole", reflect.TypeOf((*MockDatabaseRoleUpdater)(nil).UpdateDatabaseRole), arg0, arg1, arg2)
}

// MockDatabaseRoleDescriber is a mock of DatabaseRoleDescriber interface.
type MockDatabaseRoleDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseRoleDescriberMockRecorder
}

// MockDatabaseRoleDescriberMockRecorder is the mock recorder for MockDatabaseRoleDescriber.
type MockDatabaseRoleDescriberMockRecorder struct {
	mock *MockDatabaseRoleDescriber
}

// NewMockDatabaseRoleDescriber creates a new mock instance.
func NewMockDatabaseRoleDescriber(ctrl *gomock.Controller) *MockDatabaseRoleDescriber {
	mock := &MockDatabaseRoleDescriber{ctrl: ctrl}
	mock.recorder = &MockDatabaseRoleDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaseRoleDescriber) EXPECT() *MockDatabaseRoleDescriberMockRecorder {
	return m.recorder
}

// DatabaseRole mocks base method.
func (m *MockDatabaseRoleDescriber) DatabaseRole(arg0, arg1 string) (*admin.UserCustomDBRole, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseRole", arg0, arg1)
	ret0, _ := ret[0].(*admin.UserCustomDBRole)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DatabaseRole indicates an expected call of DatabaseRole.
func (mr *MockDatabaseRoleDescriberMockRecorder) DatabaseRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseRole", reflect.TypeOf((*MockDatabaseRoleDescriber)(nil).DatabaseRole), arg0, arg1)
}
