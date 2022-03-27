package service

import (
	err "driver-location-api/error"
	"driver-location-api/model"
	"driver-location-api/model/core"
	"driver-location-api/model/dto"
	"driver-location-api/repository"
	"driver-location-api/util"
	"fmt"
)

type DriverLocationService interface {
	SaveDriverLocation(dlr dto.DriverLocationRequest) (*dto.DriverLocationResponse, *err.Error)
	UploadDriverLocationFile() *err.Error
	GetNearestDriver(longitude, latitude float64, distance int) (*model.DriverInfo, *err.Error)
}

type driverLocationService struct {
	repo repository.DriverLocationRepo
}

func NewDriverLocationService(repo repository.DriverLocationRepo) DriverLocationService {
	return driverLocationService{repo: repo}
}

func (dls driverLocationService) SaveDriverLocation(dlr dto.DriverLocationRequest) (*dto.DriverLocationResponse, *err.Error) {
	e := dlr.Validate()
	if e != nil {
		return nil, e
	}
	dl := dlr.ToRepoModel()
	result, e := dls.repo.SaveDriverLocation(dl)

	if e != nil {
		return nil, e
	}
	return &dto.DriverLocationResponse{
		Type:      result.Location.Type,
		Longitude: fmt.Sprint(result.Location.Coordinates.([]float64)[0]),
		Latitude:  fmt.Sprint(result.Location.Coordinates.([]float64)[1]),
	}, nil

}
func (dls driverLocationService) UploadDriverLocationFile() *err.Error {
	return nil
}
func (dls driverLocationService) GetNearestDriver(longitude, latitude float64, radius int) (*model.DriverInfo, *err.Error) {
	riderLocation := core.NewPoint(longitude, latitude)
	drivers, _ := dls.repo.GetNearDriversLocation(riderLocation, radius)
	nearestDriver := (*drivers)[0]
	distance := util.CalculateDistance(riderLocation, nearestDriver.Location)

	return &model.DriverInfo{
		DriverLocation: nearestDriver,
		Distance:       distance,
	}, nil
}
