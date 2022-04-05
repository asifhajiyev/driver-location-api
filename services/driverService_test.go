package services

import (
	"driver-location-api/controllers/model/dto/request"
	"driver-location-api/domain/constants"
	"driver-location-api/domain/model"
	"driver-location-api/domain/model/core"
	err "driver-location-api/error"
	"driver-location-api/mocks"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func Test_driverService_GetNearestDriver_with_invalid_inputs(t *testing.T) {
	type args struct {
		sd request.SearchDriverRequest
	}
	tests := []struct {
		name      string
		args      args
		want      *model.RideInfo
		wantError *err.Error
	}{
		{
			"should return InvalidLocation error when longitude is not valid",
			args{sd: request.SearchDriverRequest{
				Radius: getIntPointer(5),
				Coordinates: core.Coordinate{
					Longitude: getFloatPointer(200),
					Latitude:  getFloatPointer(30),
				},
			}},
			nil,
			err.ValidationError(constants.ErrorInvalidLocation),
		},
		{
			"should return InvalidLocation error when latitude is not valid",
			args{sd: request.SearchDriverRequest{
				Radius: getIntPointer(5),
				Coordinates: core.Coordinate{
					Longitude: getFloatPointer(30),
					Latitude:  getFloatPointer(-100),
				},
			}},
			nil,
			err.ValidationError(constants.ErrorInvalidLocation),
		},
		{
			"should return InvalidLocation error when radius is not valid",
			args{sd: request.SearchDriverRequest{
				Radius: getIntPointer(-5),
				Coordinates: core.Coordinate{
					Longitude: getFloatPointer(100),
					Latitude:  getFloatPointer(90),
				},
			}},
			nil,
			err.ValidationError(constants.ErrorInvalidLocation),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := driverService{}
			got, gotError := ds.GetNearestDriver(tt.args.sd)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNearestDriver() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(gotError, tt.wantError) {
				t.Errorf("GetNearestDriver() gotError = %v, want %v", gotError, tt.wantError)
			}
		})
	}
}

func Test_driverService_GetNearestDriver_should_return_CouldNotGetDriverData(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDriverRepo := mocks.NewMockDriverRepository(mockCtrl)
	testDriverService := &driverService{repo: mockDriverRepo}

	longitude := getFloatPointer(-73.9667)
	latitude := getFloatPointer(40.78)
	radius := getIntPointer(5)

	errCouldNotGetDriverData := err.ServerError(constants.ErrorCouldNotGetDriverData)

	searchDriverRequest := request.SearchDriverRequest{
		Radius: radius,
		Coordinates: core.Coordinate{
			Longitude: longitude,
			Latitude:  latitude},
	}

	riderLocation := core.NewPoint(longitude, latitude)

	mockDriverRepo.
		EXPECT().
		GetNearDrivers(riderLocation, *radius).
		Return(nil, errCouldNotGetDriverData)

	got, gotError := testDriverService.GetNearestDriver(searchDriverRequest)
	if !reflect.ValueOf(got).IsNil() {
		t.Errorf("GetNearestDriver() got = %v, want %v", got, nil)
	}
	if !reflect.DeepEqual(gotError, errCouldNotGetDriverData) {
		t.Errorf("GetNearestDriver() gotError = %v, want %v", gotError, errCouldNotGetDriverData)
	}
}

func Test_driverService_GetNearestDriver_should_return_DriverNotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDriverRepo := mocks.NewMockDriverRepository(mockCtrl)
	testDriverService := &driverService{repo: mockDriverRepo}

	longitude := getFloatPointer(-73.9667)
	latitude := getFloatPointer(40.78)
	radius := getIntPointer(5)

	errDriverNotFound := err.NotFoundError(constants.ErrorDriverNotFound)
	var emptyDriverInfo = make([]*model.DriverInfo, 0)

	searchDriverRequest := request.SearchDriverRequest{
		Radius: radius,
		Coordinates: core.Coordinate{
			Longitude: longitude,
			Latitude:  latitude},
	}

	riderLocation := core.NewPoint(longitude, latitude)

	mockDriverRepo.
		EXPECT().
		GetNearDrivers(riderLocation, *radius).
		Return(emptyDriverInfo, nil)

	got, gotError := testDriverService.GetNearestDriver(searchDriverRequest)
	if !reflect.ValueOf(got).IsNil() {
		t.Errorf("GetNearestDriver() got = %v, want %v", got, nil)
	}
	if !reflect.DeepEqual(gotError, errDriverNotFound) {
		t.Errorf("GetNearestDriver() gotError = %v, want %v", gotError, errDriverNotFound)
	}
}

func getIntPointer(val int) *int {
	return &val
}
func getFloatPointer(val float64) *float64 {
	return &val
}
