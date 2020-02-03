// Code generated by MockGen. DO NOT EDIT.
// Source: gitprovider/service.go

// Package mock_gitprovider is a generated GoMock package.
package mock_gitprovider

import (
	gomock "github.com/golang/mock/gomock"
	glow "github.com/meinto/glow"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockService) Close(b glow.Branch) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", b)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockServiceMockRecorder) Close(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockService)(nil).Close), b)
}

// Publish mocks base method
func (m *MockService) Publish(b glow.Branch) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", b)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish
func (mr *MockServiceMockRecorder) Publish(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockService)(nil).Publish), b)
}

// DetectCICDOrigin mocks base method
func (m *MockService) DetectCICDOrigin() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DetectCICDOrigin")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DetectCICDOrigin indicates an expected call of DetectCICDOrigin
func (mr *MockServiceMockRecorder) DetectCICDOrigin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DetectCICDOrigin", reflect.TypeOf((*MockService)(nil).DetectCICDOrigin))
}

// GetCIBranch mocks base method
func (m *MockService) GetCIBranch() (glow.Branch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCIBranch")
	ret0, _ := ret[0].(glow.Branch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCIBranch indicates an expected call of GetCIBranch
func (mr *MockServiceMockRecorder) GetCIBranch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCIBranch", reflect.TypeOf((*MockService)(nil).GetCIBranch))
}
