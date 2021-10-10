// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/encryption (interfaces: AEAD)

// Package mock_encryption is a generated GoMock package.
package mock_encryption

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAEAD is a mock of AEAD interface.
type MockAEAD struct {
	ctrl     *gomock.Controller
	recorder *MockAEADMockRecorder
}

// MockAEADMockRecorder is the mock recorder for MockAEAD.
type MockAEADMockRecorder struct {
	mock *MockAEAD
}

// NewMockAEAD creates a new mock instance.
func NewMockAEAD(ctrl *gomock.Controller) *MockAEAD {
	mock := &MockAEAD{ctrl: ctrl}
	mock.recorder = &MockAEADMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAEAD) EXPECT() *MockAEADMockRecorder {
	return m.recorder
}

// Open mocks base method.
func (m *MockAEAD) Open(arg0 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockAEADMockRecorder) Open(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockAEAD)(nil).Open), arg0)
}

// Seal mocks base method.
func (m *MockAEAD) Seal(arg0 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seal", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Seal indicates an expected call of Seal.
func (mr *MockAEADMockRecorder) Seal(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seal", reflect.TypeOf((*MockAEAD)(nil).Seal), arg0)
}
