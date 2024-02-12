// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prysmaticlabs/prysm/v4/validator/client/iface (interfaces: PrysmBeaconChainClient)

// Package validator_mock is a generated GoMock package.
package validator_mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	validator "github.com/KP-Universe/prysm/v4/consensus-types/validator"
	iface "github.com/KP-Universe/prysm/v4/validator/client/iface"
)

// MockPrysmBeaconChainClient is a mock of PrysmBeaconChainClient interface.
type MockPrysmBeaconChainClient struct {
	ctrl     *gomock.Controller
	recorder *MockPrysmBeaconChainClientMockRecorder
}

// MockPrysmBeaconChainClientMockRecorder is the mock recorder for MockPrysmBeaconChainClient.
type MockPrysmBeaconChainClientMockRecorder struct {
	mock *MockPrysmBeaconChainClient
}

// NewMockPrysmBeaconChainClient creates a new mock instance.
func NewMockPrysmBeaconChainClient(ctrl *gomock.Controller) *MockPrysmBeaconChainClient {
	mock := &MockPrysmBeaconChainClient{ctrl: ctrl}
	mock.recorder = &MockPrysmBeaconChainClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrysmBeaconChainClient) EXPECT() *MockPrysmBeaconChainClientMockRecorder {
	return m.recorder
}

// GetValidatorCount mocks base method.
func (m *MockPrysmBeaconChainClient) GetValidatorCount(arg0 context.Context, arg1 string, arg2 []validator.Status) ([]iface.ValidatorCount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorCount", arg0, arg1, arg2)
	ret0, _ := ret[0].([]iface.ValidatorCount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidatorCount indicates an expected call of GetValidatorCount.
func (mr *MockPrysmBeaconChainClientMockRecorder) GetValidatorCount(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorCount", reflect.TypeOf((*MockPrysmBeaconChainClient)(nil).GetValidatorCount), arg0, arg1, arg2)
}
