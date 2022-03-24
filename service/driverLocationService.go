package service

import (
	err "driver-location-api/error"
	"driver-location-api/model/dto"
	"driver-location-api/repository"
)

type DriverLocationService interface {
	SaveDriverLocation(dlr dto.DriverLocationRequest) (*dto.DriverLocationResponse, *err.Error)
	UploadDriverLocationFile() *err.Error
	GetNearestDriver() *err.Error
}

type DriverLocationServiceImpl struct {
	dlr repository.DriverLocationRepo
}

func NewDriverLocationServiceImpl(repo repository.DriverLocationRepo) DriverLocationService {
	return &DriverLocationServiceImpl{dlr: repo}
}

func (dlc DriverLocationServiceImpl) SaveDriverLocation(dlr dto.DriverLocationRequest) (*dto.DriverLocationResponse, *err.Error) {
	return nil, nil
}
func (dlc DriverLocationServiceImpl) UploadDriverLocationFile() *err.Error {
	return nil
}
func (dlc DriverLocationServiceImpl) GetNearestDriver() *err.Error {
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
