// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dollarshaveclub/furan/lib/kafka (interfaces: EventBusProducer)

package mocks

import (
	lib "github.com/dollarshaveclub/furan/generated/lib"
	gomock "github.com/golang/mock/gomock"
)

// MockEventBusProducer is a mock of EventBusProducer interface
type MockEventBusProducer struct {
	ctrl     *gomock.Controller
	recorder *MockEventBusProducerMockRecorder
}

// MockEventBusProducerMockRecorder is the mock recorder for MockEventBusProducer
type MockEventBusProducerMockRecorder struct {
	mock *MockEventBusProducer
}

// NewMockEventBusProducer creates a new mock instance
func NewMockEventBusProducer(ctrl *gomock.Controller) *MockEventBusProducer {
	mock := &MockEventBusProducer{ctrl: ctrl}
	mock.recorder = &MockEventBusProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockEventBusProducer) EXPECT() *MockEventBusProducerMockRecorder {
	return _m.recorder
}

// PublishEvent mocks base method
func (_m *MockEventBusProducer) PublishEvent(_param0 *lib.BuildEvent) error {
	ret := _m.ctrl.Call(_m, "PublishEvent", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishEvent indicates an expected call of PublishEvent
func (_mr *MockEventBusProducerMockRecorder) PublishEvent(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PublishEvent", arg0)
}