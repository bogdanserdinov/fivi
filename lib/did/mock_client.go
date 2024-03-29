// Code generated by MockGen. DO NOT EDIT.
// Source: fivi/lib/did (interfaces: Interface)

// Package did is a generated GoMock package.
package did

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// RegisterDID mocks base method.
func (m *MockInterface) RegisterDID(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterDID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterDID indicates an expected call of RegisterDID.
func (mr *MockInterfaceMockRecorder) RegisterDID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterDID", reflect.TypeOf((*MockInterface)(nil).RegisterDID), arg0)
}

// ResolveDID mocks base method.
func (m *MockInterface) ResolveDID(arg0 string) (*DID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveDID", arg0)
	ret0, _ := ret[0].(*DID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveDID indicates an expected call of ResolveDID.
func (mr *MockInterfaceMockRecorder) ResolveDID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveDID", reflect.TypeOf((*MockInterface)(nil).ResolveDID), arg0)
}
