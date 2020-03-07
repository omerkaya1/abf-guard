// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/interfaces/bucket/store.go

// Package bucket is a generated GoMock package.
package bucket

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStorage is a mock of Storage interface
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// GetBucket mocks base method
func (m *MockStorage) GetBucket(arg0 string) (Bucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBucket", arg0)
	ret0, _ := ret[0].(Bucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucket indicates an expected call of GetBucket
func (mr *MockStorageMockRecorder) GetBucket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucket", reflect.TypeOf((*MockStorage)(nil).GetBucket), arg0)
}

// CheckBucket mocks base method
func (m *MockStorage) CheckBucket(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckBucket", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckBucket indicates an expected call of CheckBucket
func (mr *MockStorageMockRecorder) CheckBucket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckBucket", reflect.TypeOf((*MockStorage)(nil).CheckBucket), arg0)
}

// AddBucket mocks base method
func (m *MockStorage) AddBucket(arg0 string, arg1 Bucket) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddBucket", arg0, arg1)
}

// AddBucket indicates an expected call of AddBucket
func (mr *MockStorageMockRecorder) AddBucket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBucket", reflect.TypeOf((*MockStorage)(nil).AddBucket), arg0, arg1)
}

// RemoveBucket mocks base method
func (m *MockStorage) RemoveBucket(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveBucket", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveBucket indicates an expected call of RemoveBucket
func (mr *MockStorageMockRecorder) RemoveBucket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveBucket", reflect.TypeOf((*MockStorage)(nil).RemoveBucket), arg0)
}
