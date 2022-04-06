package request

import (
	"driver-location-api/domain/constants"
	"driver-location-api/domain/model"
	"driver-location-api/domain/model/core"
	err "driver-location-api/error"
)

type DriverLocationRequest struct {
	Type     string          `json:"type" validate:"required"`
	Location core.Coordinate `json:"location" validate:"required"`
}

type SearchDriverRequest struct {
	Radius      *int            `json:"radius" validate:"required,min=0"`
	Coordinates core.Coordinate `json:"coordinates" validate:"required"`
}

func (dlr DriverLocationRequest) ToDriverInfo() model.DriverInfo {
	t := dlr.Type
	longitude := dlr.Location.Longitude
	latitude := dlr.Location.Latitude

	return model.DriverInfo{
		Location: core.Location{
			Type:        t,
			Coordinates: []*float64{longitude, latitude},
		},
	}
}

func (dlr DriverLocationRequest) ValidateDriverLocationRequest() *err.Error {
	if !model.IsValidLongitude(dlr.Location.Longitude) ||
		!model.IsValidLatitude(dlr.Location.Latitude) ||
		!model.IsPointType(dlr.Type) {
		return err.ValidationError(constants.ErrorBadRequest, constants.ErrorInvalidLocation)
	}
	return nil
}

func (sdr SearchDriverRequest) ValidateSearchDriverRequest() *err.Error {
	if !model.IsValidLongitude(sdr.Coordinates.Longitude) ||
		!model.IsValidLatitude(sdr.Coordinates.Latitude) ||
		*sdr.Radius < 0 {
		return err.ValidationError(constants.ErrorBadRequest, constants.ErrorInvalidSearchRequest)
	}
	return nil
}
