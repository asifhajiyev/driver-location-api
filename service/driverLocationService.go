package service

import (
	err "driver-location-api/error"
	"driver-location-api/model/dto"
	"driver-location-api/repository"
	"fmt"
)

type DriverLocationService interface {
	SaveDriverLocation(dlr dto.DriverLocationRequest) (*dto.DriverLocationResponse, *err.Error)
	UploadDriverLocationFile() *err.Error
	GetNearestDriver() *err.Error
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
		fmt.Println("in service", *e)
		return nil, e
	}
	dl := dlr.ToRepoModel()
	result, e := dls.repo.SaveDriverLocation(dl)

	if e != nil {
		return nil, e
	}
	return &dto.DriverLocationResponse{
		Longitude: fmt.Sprint(result.Coordinates[0]),
		Latitude:  fmt.Sprint(dl.Coordinates[1]),
	}, nil

}
func (dls driverLocationService) UploadDriverLocationFile() *err.Error {
	return nil
}
func (dls driverLocationService) GetNearestDriver() *err.Error {
	return nil
}

/*func (dlc DriverLocationServiceImpl) SaveDriverLocation(dlr dto.DriverLocationRequest) (*dto.DriverLocationResponse, *err.Error) {
	e := dlr.Validate()
	if e != nil {
		return nil, e
	}

	dl := dlr.ToRepoModel()
	result, e := dlc.repo.SaveDriverLocation(dl)

	if e != nil {
		return nil, e
	}
	return &dto.DriverLocationResponse{
		Longitude: fmt.Sprint(result.Coordinates[0]),
		Latitude:  fmt.Sprint(dl.Coordinates[1]),
	}, nil
}*/
