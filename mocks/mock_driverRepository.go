// Code generated by MockGen. DO NOT EDIT.
// Source: C:/Users/HaciyevAB/GolandProjects/DriverLocationAPI/domain/repository/driverRepository.go

// Package mock_repository is a generated GoMock package.
package mocks

import (
	model "driver-location-api/domain/model"
	core "driver-location-api/domain/model/core"
	error "driver-location-api/error"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDriverRepository is a mock of DriverRepository interface.
type MockDriverRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDriverRepositoryMockRecorder
}

// MockDriverRepositoryMockRecorder is the mock recorder for MockDriverRepository.
type MockDriverRepositoryMockRecorder struct {
	mock *MockDriverRepository
}

// NewMockDriverRepository creates a new mock instance.
func NewMockDriverRepository(ctrl *gomock.Controller) *MockDriverRepository {
	mock := &MockDriverRepository{ctrl: ctrl}
	mock.recorder = &MockDriverRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDriverRepository) EXPECT() *MockDriverRepositoryMockRecorder {
	return m.recorder
}

// GetNearDrivers mocks base method.
func (m *MockDriverRepository) GetNearDrivers(location core.Location, radius int) ([]*model.DriverInfo, *error.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNearDrivers", location, radius)
	ret0, _ := ret[0].([]*model.DriverInfo)
	ret1, _ := ret[1].(*error.Error)
	return ret0, ret1
}

// GetNearDrivers indicates an expected call of GetNearDrivers.
func (mr *MockDriverRepositoryMockRecorder) GetNearDrivers(location, radius interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNearDrivers", reflect.TypeOf((*MockDriverRepository)(nil).GetNearDrivers), location, radius)
}

// SaveDriverInfo mocks base method.
func (m *MockDriverRepository) SaveDriverInfo(di model.DriverInfo) (*model.DriverInfo, *error.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveDriverInfo", di)
	ret0, _ := ret[0].(*model.DriverInfo)
	ret1, _ := ret[1].(*error.Error)
	return ret0, ret1
}

// SaveDriverInfo indicates an expected call of SaveDriverInfo.
func (mr *MockDriverRepositoryMockRecorder) SaveDriverInfo(di interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveDriverInfo", reflect.TypeOf((*MockDriverRepository)(nil).SaveDriverInfo), di)
}

// SaveDriverInfoSlice mocks base method.
func (m *MockDriverRepository) SaveDriverInfoSlice(di []model.DriverInfo) *error.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveDriverInfoSlice", di)
	ret0, _ := ret[0].(*error.Error)
	return ret0
}

// SaveDriverInfoSlice indicates an expected call of SaveDriverInfoSlice.
func (mr *MockDriverRepositoryMockRecorder) SaveDriverInfoSlice(di interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveDriverInfoSlice", reflect.TypeOf((*MockDriverRepository)(nil).SaveDriverInfoSlice), di)
}
