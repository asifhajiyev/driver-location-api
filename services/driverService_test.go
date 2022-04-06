package services

import (
	"driver-location-api/controllers/model/dto/request"
	"driver-location-api/controllers/model/dto/response"
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
	//arrange
	type args struct {
		sd request.SearchDriverRequest
	}
	ds := driverService{}

	tests := []struct {
		name      string
		args      args
		want      *model.RideInfo
		wantError *err.Error
	}{
		{
			"should return InvalidSearchRequest error when longitude is not valid",
			args{sd: request.SearchDriverRequest{
				Radius: getIntPointer(5),
				Coordinates: core.Coordinate{
					Longitude: getFloatPointer(200),
					Latitude:  getFloatPointer(30),
				},
			}},
			nil,
			err.ValidationError(constants.ErrorBadRequest, constants.ErrorInvalidSearchRequest),
		},
		{
			"should return InvalidSearchRequest error when latitude is not valid",
			args{sd: request.SearchDriverRequest{
				Radius: getIntPointer(5),
				Coordinates: core.Coordinate{
					Longitude: getFloatPointer(30),
					Latitude:  getFloatPointer(-100),
				},
			}},
			nil,
			err.ValidationError(constants.ErrorBadRequest, constants.ErrorInvalidSearchRequest),
		},
		{
			"should return InvalidSearchRequest error when radius is not valid",
			args{sd: request.SearchDriverRequest{
				Radius: getIntPointer(-5),
				Coordinates: core.Coordinate{
					Longitude: getFloatPointer(100),
					Latitude:  getFloatPointer(90),
				},
			}},
			nil,
			err.ValidationError(constants.ErrorBadRequest, constants.ErrorInvalidSearchRequest),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//act
			got, gotError := ds.GetNearestDriver(tt.args.sd)

			//assert
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNearestDriver() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(gotError, tt.wantError) {
				t.Errorf("GetNearestDriver() gotError = %v, want %v", gotError, tt.wantError)
			}
		})
	}
}

func Test_driverService_GetNearestDriver_should_return_error_CouldNotGetDriverData(t *testing.T) {
	//arrange
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

	//act
	got, gotError := testDriverService.GetNearestDriver(searchDriverRequest)

	//assert
	if !reflect.ValueOf(got).IsNil() {
		t.Errorf("GetNearestDriver() got = %v, want %v", got, nil)
	}
	if !reflect.DeepEqual(gotError, errCouldNotGetDriverData) {
		t.Errorf("GetNearestDriver() gotError = %v, want %v", gotError, errCouldNotGetDriverData)
	}
}

func Test_driverService_GetNearestDriver_should_return_error_DriverNotFound(t *testing.T) {
	//arrange
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

	//act
	got, gotError := testDriverService.GetNearestDriver(searchDriverRequest)

	//assert
	if !reflect.ValueOf(got).IsNil() {
		t.Errorf("GetNearestDriver() got = %v, want %v", got, nil)
	}
	if !reflect.DeepEqual(gotError, errDriverNotFound) {
		t.Errorf("GetNearestDriver() gotError = %v, want %v", gotError, errDriverNotFound)
	}
}

func Test_driverService_GetNearestDriver_should_return_DriverInfoAndDistanceAsRideInfo(t *testing.T) {
	//arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDriverRepo := mocks.NewMockDriverRepository(mockCtrl)
	testDriverService := &driverService{repo: mockDriverRepo}

	longitude := getFloatPointer(-73.9667)
	latitude := getFloatPointer(40.78)
	radius := getIntPointer(9750000)

	var driverInfo = []*model.DriverInfo{{
		Location: core.Location{
			Type:        "Point",
			Coordinates: []*float64{getFloatPointer(40.62189228), getFloatPointer(30.04352028)},
		},
	},
	}
	var rideInfo = &model.RideInfo{
		DriverInfo: *driverInfo[0],
		Distance:   9661.68,
	}

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
		Return(driverInfo, nil)

	//act
	got, gotError := testDriverService.GetNearestDriver(searchDriverRequest)

	//assert
	if !reflect.DeepEqual(got, rideInfo) {
		t.Errorf("GetNearestDriver() got = %v, want %v", got, rideInfo)
	}
	if !reflect.ValueOf(gotError).IsNil() {
		t.Errorf("GetNearestDriver() gotError = %v, want %v", gotError, nil)
	}
}

func getIntPointer(val int) *int {
	return &val
}

func getFloatPointer(val float64) *float64 {
	return &val
}

func Test_driverService_SaveDriverLocation_with_invalid_inputs(t *testing.T) {
	//arrange
	type args struct {
		dlr request.DriverLocationRequest
	}
	ds := driverService{}

	tests := []struct {
		name      string
		args      args
		want      *response.DriverLocationResponse
		wantError *err.Error
	}{
		{
			name: "should return InvalidLocation error when longitude is not valid",
			args: args{dlr: request.DriverLocationRequest{
				Type: "Point",
				Location: core.Coordinate{
					Longitude: getFloatPointer(200),
					Latitude:  getFloatPointer(30),
				},
			}},
			want:      nil,
			wantError: err.ValidationError(constants.ErrorBadRequest, constants.ErrorInvalidLocation),
		},
		{
			name: "should return InvalidLocation error when latitude is not valid",
			args: args{dlr: request.DriverLocationRequest{
				Type: "Point",
				Location: core.Coordinate{
					Longitude: getFloatPointer(150),
					Latitude:  getFloatPointer(110),
				},
			}},
			want:      nil,
			wantError: err.ValidationError(constants.ErrorBadRequest, constants.ErrorInvalidLocation),
		},
		{
			name: "should return InvalidLocation error when location type is not Point",
			args: args{dlr: request.DriverLocationRequest{
				Type: "NotPoint",
				Location: core.Coordinate{
					Longitude: getFloatPointer(150),
					Latitude:  getFloatPointer(80),
				},
			}},
			want:      nil,
			wantError: err.ValidationError(constants.ErrorBadRequest, constants.ErrorInvalidLocation),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//acct
			got, got1 := ds.SaveDriverLocation(tt.args.dlr)

			//assert
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SaveDriverLocation() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.wantError) {
				t.Errorf("SaveDriverLocation() got1 = %v, want %v", got1, tt.wantError)
			}
		})
	}
}

func Test_driverService_SaveDriverLocation_should_return_error_data_not_saved(t *testing.T) {
	//arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDriverRepo := mocks.NewMockDriverRepository(mockCtrl)
	testDriverService := &driverService{repo: mockDriverRepo}

	locationType := "Point"
	longitude := getFloatPointer(150)
	latitude := getFloatPointer(30)
	errorDataNotSaved := err.ServerError(constants.ErrorDataNotSaved)

	driverLocationRequest := request.DriverLocationRequest{
		Type: locationType,
		Location: core.Coordinate{
			Longitude: longitude,
			Latitude:  latitude,
		},
	}

	driverInfo := model.DriverInfo{
		Location: core.Location{
			Type:        locationType,
			Coordinates: []*float64{longitude, latitude},
		},
	}

	mockDriverRepo.
		EXPECT().
		SaveDriverLocation(driverInfo).
		Return(nil, errorDataNotSaved)

	//act
	got, gotError := testDriverService.SaveDriverLocation(driverLocationRequest)

	//assert
	if !reflect.ValueOf(got).IsNil() {
		t.Errorf("SaveDriverLocation() got = %v, want %v", got, nil)
	}
	if !reflect.DeepEqual(gotError, errorDataNotSaved) {
		t.Errorf("SaveDriverLocation() got1 = %v, want %v", gotError, errorDataNotSaved)
	}
}

func Test_driverService_SaveDriverLocation_should_return_saved_data(t *testing.T) {
	//arrange
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDriverRepo := mocks.NewMockDriverRepository(mockCtrl)
	testDriverService := &driverService{repo: mockDriverRepo}

	locationType := "Point"
	longitude := getFloatPointer(150)
	latitude := getFloatPointer(30)

	driverLocationRequest := request.DriverLocationRequest{
		Type: locationType,
		Location: core.Coordinate{
			Longitude: longitude,
			Latitude:  latitude,
		},
	}

	driverInfo := model.DriverInfo{
		Location: core.Location{
			Type:        locationType,
			Coordinates: []*float64{longitude, latitude},
		},
	}

	saveResult := &response.DriverLocationResponse{
		Type: locationType,
		Location: core.Coordinate{
			Longitude: longitude,
			Latitude:  latitude,
		},
	}

	mockDriverRepo.
		EXPECT().
		SaveDriverLocation(driverInfo).
		Return(&driverInfo, nil)

	//act
	got, gotError := testDriverService.SaveDriverLocation(driverLocationRequest)

	//assert
	if !reflect.DeepEqual(got, saveResult) {
		t.Errorf("SaveDriverLocation() got = %v, want %v", got, nil)
	}
	if !reflect.ValueOf(gotError).IsNil() {
		t.Errorf("SaveDriverLocation() got1 = %v, want %v", gotError, nil)
	}
}
