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
	GetNearestDriver(longitude, latitude float64, distance int) (*model.RideInfo, *err.Error)
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
	di := dlr.ToRepoModel()
	result, e := dls.repo.SaveDriverLocation(di)

	if e != nil {
		return nil, e
	}
	return &dto.DriverLocationResponse{
		Type: result.Location.Type,
		Location: core.Coordinate{
			Longitude: result.Location.Coordinates[0],
			Latitude:  result.Location.Coordinates[1],
		},
	}, nil
}

func (dls driverLocationService) UploadDriverLocationFile() *err.Error {
	return nil
}

func (dls driverLocationService) GetNearestDriver(longitude, latitude float64, radius int) (*model.RideInfo, *err.Error) {
	if !model.IsValidLongitude(longitude) || !model.IsValidLatitude(latitude) {
		return nil, err.ValidationError("longitude and latitude should be in the right range")
	}

	riderLocation := core.NewPoint(longitude, latitude)
	drivers, _ := dls.repo.GetNearDriversInfo(riderLocation, radius)
	fmt.Println("drivers size", len(drivers))
	nearestDriver := drivers[0]

	riderCoordinates := core.GetCoordinates(riderLocation)
	nearestDriverCoordinates := core.GetCoordinates(nearestDriver.Location)

	calculator := util.CalculateByHaversine{}
	distance := calculator.Calculate(riderCoordinates, nearestDriverCoordinates)

	return &model.RideInfo{
		DriverInfo: *nearestDriver,
		Distance:   distance,
	}, nil
}
