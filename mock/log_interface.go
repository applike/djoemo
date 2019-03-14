// Code generated by MockGen. DO NOT EDIT.
// Source: log_interface.go

// Package mock_dynamo is a generated GoMock package.
package mock

import (
	dynamo "adjoe.io/djoemo"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLogInterface is a mock of LogInterface interface
type MockLogInterface struct {
	ctrl     *gomock.Controller
	recorder *MockLogInterfaceMockRecorder
}

// MockLogInterfaceMockRecorder is the mock recorder for MockLogInterface
type MockLogInterfaceMockRecorder struct {
	mock *MockLogInterface
}

// NewMockLogInterface creates a new mock instance
func NewMockLogInterface(ctrl *gomock.Controller) *MockLogInterface {
	mock := &MockLogInterface{ctrl: ctrl}
	mock.recorder = &MockLogInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogInterface) EXPECT() *MockLogInterfaceMockRecorder {
	return m.recorder
}

// WithContext mocks base method
func (m *MockLogInterface) WithContext(ctx context.Context) dynamo.LogInterface {
	ret := m.ctrl.Call(m, "WithContext", ctx)
	ret0, _ := ret[0].(dynamo.LogInterface)
	return ret0
}

// WithContext indicates an expected call of WithContext
func (mr *MockLogInterfaceMockRecorder) WithContext(ctx interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContext", reflect.TypeOf((*MockLogInterface)(nil).WithContext), ctx)
}

// WithFields mocks base method
func (m *MockLogInterface) WithFields(fields map[string]interface{}) dynamo.LogInterface {
	ret := m.ctrl.Call(m, "WithFields", fields)
	ret0, _ := ret[0].(dynamo.LogInterface)
	return ret0
}

// WithFields indicates an expected call of WithFields
func (mr *MockLogInterfaceMockRecorder) WithFields(fields interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithFields", reflect.TypeOf((*MockLogInterface)(nil).WithFields), fields)
}

// Infof mocks base method
func (m *MockLogInterface) Infof(format string, args ...interface{}) {
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof
func (mr *MockLogInterfaceMockRecorder) Infof(format interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLogInterface)(nil).Infof), varargs...)
}

// Warnf mocks base method
func (m *MockLogInterface) Warnf(format string, args ...interface{}) {
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf
func (mr *MockLogInterfaceMockRecorder) Warnf(format interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*MockLogInterface)(nil).Warnf), varargs...)
}

// Errorf mocks base method
func (m *MockLogInterface) Errorf(format string, args ...interface{}) {
	varargs := []interface{}{format}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf
func (mr *MockLogInterfaceMockRecorder) Errorf(format interface{}, args ...interface{}) *gomock.Call {
	varargs := append([]interface{}{format}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogInterface)(nil).Errorf), varargs...)
}
