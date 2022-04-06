package repository

import (
	"driver-location-api/domain/model"
	"driver-location-api/domain/model/core"
	err "driver-location-api/error"
)

type DriverRepository interface {
	SaveDriverInfo(di model.DriverInfo) (*model.DriverInfo, *err.Error)
	SaveDriverInfoSlice(di []model.DriverInfo) *err.Error
	GetNearDrivers(location core.Location, radius int) ([]*model.DriverInfo, *err.Error)
}
