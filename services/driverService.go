package services

import (
	"driver-location-api/controllers/model/dto/request"
	"driver-location-api/controllers/model/dto/response"
	"driver-location-api/domain/constants"
	"driver-location-api/domain/model"
	"driver-location-api/domain/model/core"
	"driver-location-api/domain/repository"
	err "driver-location-api/error"
	"driver-location-api/logger"
	"driver-location-api/util"
	"mime/multipart"
	"os"
)

type DriverService interface {
	SaveDriverLocation(dlr request.DriverLocationRequest) (*response.DriverLocationResponse, *err.Error)
	SaveDriverLocationFile(fh *multipart.FileHeader) (string, *err.Error)
	GetNearestDriver(sd request.SearchDriverRequest) (*model.RideInfo, *err.Error)
}

type driverService struct {
	repo repository.DriverRepository
}

func NewDriverService(repo repository.DriverRepository) DriverService {
	return driverService{repo: repo}
}

func (ds driverService) SaveDriverLocation(dlr request.DriverLocationRequest) (*response.DriverLocationResponse, *err.Error) {
	logger.Info("SaveDriverLocation.begin")
	e := dlr.ValidateDriverLocationRequest()
	if e != nil {
		logger.Error("SaveDriverLocation.error", e)
		return nil, e
	}
	di := dlr.ToDriverInfo()
	result, e := ds.repo.SaveDriverInfo(di)

	if e != nil {
		logger.Error("SaveDriverLocation.error", e)
		return nil, e
	}

	resp := &response.DriverLocationResponse{
		Type: result.Location.Type,
		Location: core.Coordinate{
			Longitude: result.Location.Coordinates[0],
			Latitude:  result.Location.Coordinates[1],
		},
	}
	logger.Info("SaveDriverLocation.end", resp)
	return resp, nil
}

func (ds driverService) GetNearestDriver(sd request.SearchDriverRequest) (*model.RideInfo, *err.Error) {
	logger.Info("GetNearestDriver.begin")
	longitude := sd.Coordinates.Longitude
	latitude := sd.Coordinates.Latitude
	radius := sd.Radius

	if e := sd.ValidateSearchDriverRequest(); e != nil {
		return nil, e
	}

	riderLocation := core.NewPoint(longitude, latitude)
	drivers, er := ds.repo.GetNearDrivers(riderLocation, *radius)

	if er != nil {
		return nil, er
	}

	if len(drivers) == 0 {
		return nil, err.NotFoundError(constants.ErrorDriverNotFound)
	}
	nearestDriver := drivers[0]
	distance := calculateDistance(riderLocation, nearestDriver.Location)

	rideInfo := &model.RideInfo{
		DriverInfo: *nearestDriver,
		Distance:   distance,
	}

	logger.Info("GetNearestDriver.end", rideInfo)
	return rideInfo, nil
}

func (ds driverService) SaveDriverLocationFile(fh *multipart.FileHeader) (string, *err.Error) {
	logger.Info("SaveDriverLocationFile.begin")
	content := util.CsvToSlice(fh)

	var dlUploadPatchSize = util.StringToInt(os.Getenv("bitaksi_task_INSERT_DOC_NUM_AT_ONCE"))
	patchData := make([][]string, 0)

	for _, v := range content {
		patchData = append(patchData, v)
		if len(patchData) == dlUploadPatchSize {
			driverInfoSlice := toDriverInfoSlice(patchData)
			go ds.repo.SaveDriverInfoSlice(driverInfoSlice)
			patchData = nil
		}
	}
	if len(patchData) > 0 {
		driverInfoSlice := toDriverInfoSlice(patchData)
		go ds.repo.SaveDriverInfoSlice(driverInfoSlice)
	}
	logger.Info("SaveDriverLocationFile.end")
	return constants.SavingDriverData, nil
}

func toDriverInfoSlice(s [][]string) []model.DriverInfo {
	var dis []model.DriverInfo

	for i := 0; i < len(s); i++ {
		longitude := util.StringToFloat(s[i][0])
		latitude := util.StringToFloat(s[i][1])
		di := model.DriverInfo{Location: core.Location{
			Type:        constants.LocationTypePoint,
			Coordinates: []*float64{&longitude, &latitude},
		}}
		dis = append(dis, di)
	}

	return dis
}

func calculateDistance(from core.Location, to core.Location) float64 {
	riderCoordinates := core.GetCoordinates(from)
	nearestDriverCoordinates := core.GetCoordinates(to)

	calculator := util.CalculateByHaversine{}
	distance := calculator.Calculate(riderCoordinates, nearestDriverCoordinates)

	return distance
}
