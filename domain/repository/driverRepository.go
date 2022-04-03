package repository

import (
	"driver-location-api/domain/model"
	"driver-location-api/domain/model/core"
	err "driver-location-api/error"
)

type DriverRepository interface {
	SaveDriverLocation(di model.DriverInfo) (*model.DriverInfo, *err.Error)
	SaveDriverLocationFile(di []model.DriverInfo) *err.Error
	GetNearDrivers(g core.Location, radius int) ([]*model.DriverInfo, *err.Error)
}
