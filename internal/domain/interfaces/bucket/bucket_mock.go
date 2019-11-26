// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/interfaces/bucket/bucket.go

// Package bucket is a generated GoMock package.
package bucket

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockBucket is a mock of Bucket interface
type MockBucket struct {
	ctrl     *gomock.Controller
	recorder *MockBucketMockRecorder
}

// MockBucketMockRecorder is the mock recorder for MockBucket
type MockBucketMockRecorder struct {
	mock *MockBucket
}

// NewMockBucket creates a new mock instance
func NewMockBucket(ctrl *gomock.Controller) *MockBucket {
	mock := &MockBucket{ctrl: ctrl}
	mock.recorder = &MockBucketMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBucket) EXPECT() *MockBucketMockRecorder {
	return m.recorder
}

// Decrement mocks base method
func (m *MockBucket) Decrement() bool {
	ret := m.ctrl.Call(m, "Decrement")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Decrement indicates an expected call of Decrement
func (mr *MockBucketMockRecorder) Decrement() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decrement", reflect.TypeOf((*MockBucket)(nil).Decrement))
}

// Stop mocks base method
func (m *MockBucket) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockBucketMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockBucket)(nil).Stop))
}
