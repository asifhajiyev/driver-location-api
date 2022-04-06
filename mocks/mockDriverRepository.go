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

// SaveDriverLocation mocks base method.
func (m *MockDriverRepository) SaveDriverLocation(di model.DriverInfo) (*model.DriverInfo, *error.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveDriverLocation", di)
	ret0, _ := ret[0].(*model.DriverInfo)
	ret1, _ := ret[1].(*error.Error)
	return ret0, ret1
}

// SaveDriverLocation indicates an expected call of SaveDriverLocation.
func (mr *MockDriverRepositoryMockRecorder) SaveDriverLocation(di interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveDriverLocation", reflect.TypeOf((*MockDriverRepository)(nil).SaveDriverLocation), di)
}

// SaveDriverLocationFile mocks base method.
func (m *MockDriverRepository) SaveDriverLocationFile(di []model.DriverInfo) *error.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveDriverLocationFile", di)
	ret0, _ := ret[0].(*error.Error)
	return ret0
}

// SaveDriverLocationFile indicates an expected call of SaveDriverLocationFile.
func (mr *MockDriverRepositoryMockRecorder) SaveDriverLocationFile(di interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveDriverLocationFile", reflect.TypeOf((*MockDriverRepository)(nil).SaveDriverLocationFile), di)
}