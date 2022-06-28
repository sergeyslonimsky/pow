// Code generated by MockGen. DO NOT EDIT.
// Source: ./quote.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockquoteStorage is a mock of quoteStorage interface.
type MockquoteStorage struct {
	ctrl     *gomock.Controller
	recorder *MockquoteStorageMockRecorder
}

// MockquoteStorageMockRecorder is the mock recorder for MockquoteStorage.
type MockquoteStorageMockRecorder struct {
	mock *MockquoteStorage
}

// NewMockquoteStorage creates a new mock instance.
func NewMockquoteStorage(ctrl *gomock.Controller) *MockquoteStorage {
	mock := &MockquoteStorage{ctrl: ctrl}
	mock.recorder = &MockquoteStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockquoteStorage) EXPECT() *MockquoteStorageMockRecorder {
	return m.recorder
}

// GetRandomQuote mocks base method.
func (m *MockquoteStorage) GetRandomQuote() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandomQuote")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRandomQuote indicates an expected call of GetRandomQuote.
func (mr *MockquoteStorageMockRecorder) GetRandomQuote() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandomQuote", reflect.TypeOf((*MockquoteStorage)(nil).GetRandomQuote))
}

// Mockpow is a mock of pow interface.
type Mockpow struct {
	ctrl     *gomock.Controller
	recorder *MockpowMockRecorder
}

// MockpowMockRecorder is the mock recorder for Mockpow.
type MockpowMockRecorder struct {
	mock *Mockpow
}

// NewMockpow creates a new mock instance.
func NewMockpow(ctrl *gomock.Controller) *Mockpow {
	mock := &Mockpow{ctrl: ctrl}
	mock.recorder = &MockpowMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockpow) EXPECT() *MockpowMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *Mockpow) Generate(ctx context.Context, resource string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate", ctx, resource)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockpowMockRecorder) Generate(ctx, resource interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*Mockpow)(nil).Generate), ctx, resource)
}

// Verify mocks base method.
func (m *Mockpow) Verify(ctx context.Context, resource, data string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", ctx, resource, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Verify indicates an expected call of Verify.
func (mr *MockpowMockRecorder) Verify(ctx, resource, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*Mockpow)(nil).Verify), ctx, resource, data)
}