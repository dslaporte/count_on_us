// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/entity/account/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "count_on_us/internal/entity/account"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAccountRepositoryInterface is a mock of AccountRepositoryInterface interface.
type MockAccountRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositoryInterfaceMockRecorder
}

// MockAccountRepositoryInterfaceMockRecorder is the mock recorder for MockAccountRepositoryInterface.
type MockAccountRepositoryInterfaceMockRecorder struct {
	mock *MockAccountRepositoryInterface
}

// NewMockAccountRepositoryInterface creates a new mock instance.
func NewMockAccountRepositoryInterface(ctrl *gomock.Controller) *MockAccountRepositoryInterface {
	mock := &MockAccountRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockAccountRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepositoryInterface) EXPECT() *MockAccountRepositoryInterfaceMockRecorder {
	return m.recorder
}

// FindByID mocks base method.
func (m *MockAccountRepositoryInterface) FindByID(id string) (*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", id)
	ret0, _ := ret[0].(*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockAccountRepositoryInterfaceMockRecorder) FindByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockAccountRepositoryInterface)(nil).FindByID), id)
}

// List mocks base method.
func (m *MockAccountRepositoryInterface) List(limit, offset int) ([]*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", limit, offset)
	ret0, _ := ret[0].([]*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAccountRepositoryInterfaceMockRecorder) List(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAccountRepositoryInterface)(nil).List), limit, offset)
}

// Save mocks base method.
func (m *MockAccountRepositoryInterface) Save(account *entity.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", account)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockAccountRepositoryInterfaceMockRecorder) Save(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockAccountRepositoryInterface)(nil).Save), account)
}

// Update mocks base method.
func (m *MockAccountRepositoryInterface) Update(account *entity.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", account)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAccountRepositoryInterfaceMockRecorder) Update(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAccountRepositoryInterface)(nil).Update), account)
}