// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prysmaticlabs/prysm/v4/validator/client/beacon-api (interfaces: GenesisProvider)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	beacon "github.com/KP-Universe/prysm/v4/beacon-chain/rpc/eth/beacon"
)

// MockGenesisProvider is a mock of GenesisProvider interface.
type MockGenesisProvider struct {
	ctrl     *gomock.Controller
	recorder *MockGenesisProviderMockRecorder
}

// MockGenesisProviderMockRecorder is the mock recorder for MockGenesisProvider.
type MockGenesisProviderMockRecorder struct {
	mock *MockGenesisProvider
}

// NewMockGenesisProvider creates a new mock instance.
func NewMockGenesisProvider(ctrl *gomock.Controller) *MockGenesisProvider {
	mock := &MockGenesisProvider{ctrl: ctrl}
	mock.recorder = &MockGenesisProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGenesisProvider) EXPECT() *MockGenesisProviderMockRecorder {
	return m.recorder
}

// GetGenesis mocks base method.
func (m *MockGenesisProvider) GetGenesis(arg0 context.Context) (*beacon.Genesis, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenesis", arg0)
	ret0, _ := ret[0].(*beacon.Genesis)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGenesis indicates an expected call of GetGenesis.
func (mr *MockGenesisProviderMockRecorder) GetGenesis(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenesis", reflect.TypeOf((*MockGenesisProvider)(nil).GetGenesis), arg0)
}
