// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dollarshaveclub/furan/lib/consul (interfaces: ConsulKV)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	api "github.com/hashicorp/consul/api"
	reflect "reflect"
)

// MockConsulKV is a mock of ConsulKV interface
type MockConsulKV struct {
	ctrl     *gomock.Controller
	recorder *MockConsulKVMockRecorder
}

// MockConsulKVMockRecorder is the mock recorder for MockConsulKV
type MockConsulKVMockRecorder struct {
	mock *MockConsulKV
}

// NewMockConsulKV creates a new mock instance
func NewMockConsulKV(ctrl *gomock.Controller) *MockConsulKV {
	mock := &MockConsulKV{ctrl: ctrl}
	mock.recorder = &MockConsulKVMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConsulKV) EXPECT() *MockConsulKVMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockConsulKV) Delete(arg0 string, arg1 *api.WriteOptions) (*api.WriteMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(*api.WriteMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockConsulKVMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockConsulKV)(nil).Delete), arg0, arg1)
}

// Get mocks base method
func (m *MockConsulKV) Get(arg0 string, arg1 *api.QueryOptions) (*api.KVPair, *api.QueryMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*api.KVPair)
	ret1, _ := ret[1].(*api.QueryMeta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get
func (mr *MockConsulKVMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockConsulKV)(nil).Get), arg0, arg1)
}

// Keys mocks base method
func (m *MockConsulKV) Keys(arg0, arg1 string, arg2 *api.QueryOptions) ([]string, *api.QueryMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Keys", arg0, arg1, arg2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(*api.QueryMeta)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Keys indicates an expected call of Keys
func (mr *MockConsulKVMockRecorder) Keys(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Keys", reflect.TypeOf((*MockConsulKV)(nil).Keys), arg0, arg1, arg2)
}

// Put mocks base method
func (m *MockConsulKV) Put(arg0 *api.KVPair, arg1 *api.WriteOptions) (*api.WriteMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", arg0, arg1)
	ret0, _ := ret[0].(*api.WriteMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put
func (mr *MockConsulKVMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockConsulKV)(nil).Put), arg0, arg1)
}
