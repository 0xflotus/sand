// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/sand/network/netmanager (interfaces: NetManager)

// Package netmanagermock is a generated GoMock package.
package netmanagermock

import (
	context "context"
	reflect "reflect"

	params "github.com/Scalingo/sand/api/params"
	types "github.com/Scalingo/sand/api/types"
	gomock "github.com/golang/mock/gomock"
)

// MockNetManager is a mock of NetManager interface
type MockNetManager struct {
	ctrl     *gomock.Controller
	recorder *MockNetManagerMockRecorder
}

// MockNetManagerMockRecorder is the mock recorder for MockNetManager
type MockNetManagerMockRecorder struct {
	mock *MockNetManager
}

// NewMockNetManager creates a new mock instance
func NewMockNetManager(ctrl *gomock.Controller) *MockNetManager {
	mock := &MockNetManager{ctrl: ctrl}
	mock.recorder = &MockNetManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetManager) EXPECT() *MockNetManagerMockRecorder {
	return m.recorder
}

// AddEndpointNeigh mocks base method
func (m *MockNetManager) AddEndpointNeigh(arg0 context.Context, arg1 types.Network, arg2 types.Endpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEndpointNeigh", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEndpointNeigh indicates an expected call of AddEndpointNeigh
func (mr *MockNetManagerMockRecorder) AddEndpointNeigh(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEndpointNeigh", reflect.TypeOf((*MockNetManager)(nil).AddEndpointNeigh), arg0, arg1, arg2)
}

// Deactivate mocks base method
func (m *MockNetManager) Deactivate(arg0 context.Context, arg1 types.Network) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deactivate", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deactivate indicates an expected call of Deactivate
func (mr *MockNetManagerMockRecorder) Deactivate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deactivate", reflect.TypeOf((*MockNetManager)(nil).Deactivate), arg0, arg1)
}

// DeleteEndpoint mocks base method
func (m *MockNetManager) DeleteEndpoint(arg0 context.Context, arg1 types.Network, arg2 types.Endpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEndpoint", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEndpoint indicates an expected call of DeleteEndpoint
func (mr *MockNetManagerMockRecorder) DeleteEndpoint(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEndpoint", reflect.TypeOf((*MockNetManager)(nil).DeleteEndpoint), arg0, arg1, arg2)
}

// Ensure mocks base method
func (m *MockNetManager) Ensure(arg0 context.Context, arg1 types.Network) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ensure", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ensure indicates an expected call of Ensure
func (mr *MockNetManagerMockRecorder) Ensure(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ensure", reflect.TypeOf((*MockNetManager)(nil).Ensure), arg0, arg1)
}

// EnsureEndpoint mocks base method
func (m *MockNetManager) EnsureEndpoint(arg0 context.Context, arg1 types.Network, arg2 types.Endpoint, arg3 params.EndpointActivate) (types.Endpoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureEndpoint", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(types.Endpoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureEndpoint indicates an expected call of EnsureEndpoint
func (mr *MockNetManagerMockRecorder) EnsureEndpoint(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureEndpoint", reflect.TypeOf((*MockNetManager)(nil).EnsureEndpoint), arg0, arg1, arg2, arg3)
}

// EnsureEndpointsNeigh mocks base method
func (m *MockNetManager) EnsureEndpointsNeigh(arg0 context.Context, arg1 types.Network, arg2 []types.Endpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureEndpointsNeigh", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnsureEndpointsNeigh indicates an expected call of EnsureEndpointsNeigh
func (mr *MockNetManagerMockRecorder) EnsureEndpointsNeigh(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureEndpointsNeigh", reflect.TypeOf((*MockNetManager)(nil).EnsureEndpointsNeigh), arg0, arg1, arg2)
}

// ListenNetworkChange mocks base method
func (m *MockNetManager) ListenNetworkChange(arg0 context.Context, arg1 types.Network) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenNetworkChange", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ListenNetworkChange indicates an expected call of ListenNetworkChange
func (mr *MockNetManagerMockRecorder) ListenNetworkChange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenNetworkChange", reflect.TypeOf((*MockNetManager)(nil).ListenNetworkChange), arg0, arg1)
}

// RemoveEndpointNeigh mocks base method
func (m *MockNetManager) RemoveEndpointNeigh(arg0 context.Context, arg1 types.Network, arg2 types.Endpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveEndpointNeigh", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveEndpointNeigh indicates an expected call of RemoveEndpointNeigh
func (mr *MockNetManagerMockRecorder) RemoveEndpointNeigh(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveEndpointNeigh", reflect.TypeOf((*MockNetManager)(nil).RemoveEndpointNeigh), arg0, arg1, arg2)
}

// StopListenNetworkChange mocks base method
func (m *MockNetManager) StopListenNetworkChange(arg0 context.Context, arg1 types.Network) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopListenNetworkChange", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopListenNetworkChange indicates an expected call of StopListenNetworkChange
func (mr *MockNetManagerMockRecorder) StopListenNetworkChange(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopListenNetworkChange", reflect.TypeOf((*MockNetManager)(nil).StopListenNetworkChange), arg0, arg1)
}
