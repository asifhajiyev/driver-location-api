package service

import (
	"driver-location-api/controllers/model/dto/request"
	"driver-location-api/controllers/model/dto/response"
	"driver-location-api/domain/model"
	"driver-location-api/domain/model/core"
	"driver-location-api/domain/repository"
	err "driver-location-api/error"
	"driver-location-api/util"
	"fmt"
	"mime/multipart"
	"os"
)

type DriverLocationService interface {
	SaveDriverLocation(dlr request.DriverLocationRequest) (*response.DriverLocationResponse, *err.Error)
	SaveDriverLocationFile(fh *multipart.FileHeader) *err.Error
	GetNearestDriver(sd request.SearchDriverRequest) (*model.RideInfo, *err.Error)
}

type driverLocationService struct {
	repo repository.DriverLocationRepo
}

func NewDriverLocationService(repo repository.DriverLocationRepo) DriverLocationService {
	return driverLocationService{repo: repo}
}

func (dls driverLocationService) SaveDriverLocation(dlr request.DriverLocationRequest) (*response.DriverLocationResponse, *err.Error) {
	e := dlr.Validate()
	if e != nil {
		return nil, e
	}
	di := dlr.ToDriverInfo()
	result, e := dls.repo.SaveDriverLocation(di)

	if e != nil {
		return nil, e
	}
	return &response.DriverLocationResponse{
		Type: result.Location.Type,
		Location: core.Coordinate{
			Longitude: result.Location.Coordinates[0],
			Latitude:  result.Location.Coordinates[1],
		},
	}, nil
}

func (dls driverLocationService) GetNearestDriver(sd request.SearchDriverRequest) (*model.RideInfo, *err.Error) {
	longitude := sd.Coordinates.Longitude
	latitude := sd.Coordinates.Latitude
	radius := sd.Radius

	if !model.IsValidLongitude(longitude) || !model.IsValidLatitude(latitude) {
		return nil, err.ValidationError("longitude and latitude should be in the right range")
	}

	riderLocation := core.NewPoint(longitude, latitude)
	drivers, _ := dls.repo.GetNearDrivers(riderLocation, radius)
	fmt.Println("drivers size", len(drivers))

	if len(drivers) == 0 {
		return nil, err.NotFoundError("no drivers found in given radius")
	}
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

func (dls driverLocationService) SaveDriverLocationFile(fh *multipart.FileHeader) *err.Error {
	content := util.CsvToSlice(fh)

	var dlUploadPatchSize = util.StringToInt(os.Getenv("bitaksi_task_INSERT_DOC_NUM_AT_ONCE"))
	patchData := make([][]string, 0)

	for _, v := range content {
		patchData = append(patchData, v)
		if len(patchData) == dlUploadPatchSize {
			fmt.Println("data in len", len(patchData))
			go toDriverInfoSliceAndUpload(dls, patchData)
			patchData = nil
		}
	}
	if len(patchData) > 0 {
		fmt.Println("data in remainder", len(patchData))
		go toDriverInfoSliceAndUpload(dls, patchData)
	}

	return nil
}

func toDriverInfoSliceAndUpload(dls driverLocationService, s [][]string) {
	var dis []model.DriverInfo

	for i := 0; i < len(s); i++ {
		longitude := util.StringToFloat(s[i][0])
		latitude := util.StringToFloat(s[i][1])
		di := model.DriverInfo{Location: core.Location{
			Type:        "Point",
			Coordinates: []float64{longitude, latitude},
		}}
		dis = append(dis, di)
	}

	dls.repo.SaveDriverLocationFile(dis)
}
