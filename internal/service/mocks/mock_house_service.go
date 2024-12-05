// Code generated by MockGen. DO NOT EDIT.
// Source: house_service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	model "houseService/internal/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHouse is a mock of House interface.
type MockHouse struct {
	ctrl     *gomock.Controller
	recorder *MockHouseMockRecorder
}

// MockHouseMockRecorder is the mock recorder for MockHouse.
type MockHouseMockRecorder struct {
	mock *MockHouse
}

// NewMockHouse creates a new mock instance.
func NewMockHouse(ctrl *gomock.Controller) *MockHouse {
	mock := &MockHouse{ctrl: ctrl}
	mock.recorder = &MockHouseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHouse) EXPECT() *MockHouseMockRecorder {
	return m.recorder
}

// HouseCreate mocks base method.
func (m *MockHouse) HouseCreate(arg0 *model.HouseCreate) (*model.House, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HouseCreate", arg0)
	ret0, _ := ret[0].(*model.House)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HouseCreate indicates an expected call of HouseCreate.
func (mr *MockHouseMockRecorder) HouseCreate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HouseCreate", reflect.TypeOf((*MockHouse)(nil).HouseCreate), arg0)
}

// HouseGetById mocks base method.
func (m *MockHouse) HouseGetById(arg0 int) (*model.House, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HouseGetById", arg0)
	ret0, _ := ret[0].(*model.House)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HouseGetById indicates an expected call of HouseGetById.
func (mr *MockHouseMockRecorder) HouseGetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HouseGetById", reflect.TypeOf((*MockHouse)(nil).HouseGetById), arg0)
}

// HouseUpdate mocks base method.
func (m *MockHouse) HouseUpdate(arg0 *model.House) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HouseUpdate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// HouseUpdate indicates an expected call of HouseUpdate.
func (mr *MockHouseMockRecorder) HouseUpdate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HouseUpdate", reflect.TypeOf((*MockHouse)(nil).HouseUpdate), arg0)
}
