// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/interfaces/bucket/manager.go

// Package bucket is a generated GoMock package.
package bucket

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// Dispatch mocks base method
func (m *MockManager) Dispatch(login, pwd, ip string) (bool, error) {
	ret := m.ctrl.Call(m, "Dispatch", login, pwd, ip)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Dispatch indicates an expected call of Dispatch
func (mr *MockManagerMockRecorder) Dispatch(login, pwd, ip interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispatch", reflect.TypeOf((*MockManager)(nil).Dispatch), login, pwd, ip)
}

// FlushBuckets mocks base method
func (m *MockManager) FlushBuckets(login, ip string) error {
	ret := m.ctrl.Call(m, "FlushBuckets", login, ip)
	ret0, _ := ret[0].(error)
	return ret0
}

// FlushBuckets indicates an expected call of FlushBuckets
func (mr *MockManagerMockRecorder) FlushBuckets(login, ip interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushBuckets", reflect.TypeOf((*MockManager)(nil).FlushBuckets), login, ip)
}

// PurgeBucket mocks base method
func (m *MockManager) PurgeBucket(name string) error {
	ret := m.ctrl.Call(m, "PurgeBucket", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// PurgeBucket indicates an expected call of PurgeBucket
func (mr *MockManagerMockRecorder) PurgeBucket(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurgeBucket", reflect.TypeOf((*MockManager)(nil).PurgeBucket), name)
}

// GetErrChan mocks base method
func (m *MockManager) GetErrChan() chan error {
	ret := m.ctrl.Call(m, "GetErrChan")
	ret0, _ := ret[0].(chan error)
	return ret0
}

// GetErrChan indicates an expected call of GetErrChan
func (mr *MockManagerMockRecorder) GetErrChan() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetErrChan", reflect.TypeOf((*MockManager)(nil).GetErrChan))
}
