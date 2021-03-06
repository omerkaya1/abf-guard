// Code generated by MockGen. DO NOT EDIT.
// Source: internal/db/storage.go

// Package db is a generated GoMock package.
package db

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStorageManager is a mock of StorageManager interface
type MockStorageManager struct {
	ctrl     *gomock.Controller
	recorder *MockStorageManagerMockRecorder
}

// MockStorageManagerMockRecorder is the mock recorder for MockStorageManager
type MockStorageManagerMockRecorder struct {
	mock *MockStorageManager
}

// NewMockStorageManager creates a new mock instance
func NewMockStorageManager(ctrl *gomock.Controller) *MockStorageManager {
	mock := &MockStorageManager{ctrl: ctrl}
	mock.recorder = &MockStorageManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorageManager) EXPECT() *MockStorageManagerMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockStorageManager) Add(ctx context.Context, ip string, blacklist bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, ip, blacklist)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockStorageManagerMockRecorder) Add(ctx, ip, blacklist interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockStorageManager)(nil).Add), ctx, ip, blacklist)
}

// Delete mocks base method
func (m *MockStorageManager) Delete(ctx context.Context, ip string, blacklist bool) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, ip, blacklist)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockStorageManagerMockRecorder) Delete(ctx, ip, blacklist interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStorageManager)(nil).Delete), ctx, ip, blacklist)
}

// GetIPList mocks base method
func (m *MockStorageManager) GetIPList(ctx context.Context, blacklist bool) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIPList", ctx, blacklist)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIPList indicates an expected call of GetIPList
func (mr *MockStorageManagerMockRecorder) GetIPList(ctx, blacklist interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPList", reflect.TypeOf((*MockStorageManager)(nil).GetIPList), ctx, blacklist)
}

// GreenLightPass mocks base method
func (m *MockStorageManager) GreenLightPass(ctx context.Context, ip string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GreenLightPass", ctx, ip)
	ret0, _ := ret[0].(error)
	return ret0
}

// GreenLightPass indicates an expected call of GreenLightPass
func (mr *MockStorageManagerMockRecorder) GreenLightPass(ctx, ip interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GreenLightPass", reflect.TypeOf((*MockStorageManager)(nil).GreenLightPass), ctx, ip)
}
