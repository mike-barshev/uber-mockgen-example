// Code generated by MockGen. DO NOT EDIT.
// Source: example_testing.go
//
// Generated by this command:
//
//	mockgen -source=example_testing.go -destination=mock/example_testing.go_mock.go -package=examples_mock
//
// Package examples_mock is a generated GoMock package.
package examples_mock

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPrefixCounter is a mock of PrefixCounter interface.
type MockPrefixCounter struct {
	ctrl     *gomock.Controller
	recorder *MockPrefixCounterMockRecorder
}

// MockPrefixCounterMockRecorder is the mock recorder for MockPrefixCounter.
type MockPrefixCounterMockRecorder struct {
	mock *MockPrefixCounter
}

// NewMockPrefixCounter creates a new mock instance.
func NewMockPrefixCounter(ctrl *gomock.Controller) *MockPrefixCounter {
	mock := &MockPrefixCounter{ctrl: ctrl}
	mock.recorder = &MockPrefixCounterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrefixCounter) EXPECT() *MockPrefixCounterMockRecorder {
	return m.recorder
}

// AddCountPostfix mocks base method.
func (m *MockPrefixCounter) AddCountPostfix(i int, postfix, str string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCountPostfix", i, postfix, str)
	ret0, _ := ret[0].(string)
	return ret0
}

// AddCountPostfix indicates an expected call of AddCountPostfix.
func (mr *MockPrefixCounterMockRecorder) AddCountPostfix(i, postfix, str any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCountPostfix", reflect.TypeOf((*MockPrefixCounter)(nil).AddCountPostfix), i, postfix, str)
}

// AddCountPrefix mocks base method.
func (m *MockPrefixCounter) AddCountPrefix(i int, prefix, str string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCountPrefix", i, prefix, str)
	ret0, _ := ret[0].(string)
	return ret0
}

// AddCountPrefix indicates an expected call of AddCountPrefix.
func (mr *MockPrefixCounterMockRecorder) AddCountPrefix(i, prefix, str any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCountPrefix", reflect.TypeOf((*MockPrefixCounter)(nil).AddCountPrefix), i, prefix, str)
}
