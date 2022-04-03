package request

import (
	"driver-location-api/domain/constants"
	"driver-location-api/domain/model"
	"driver-location-api/domain/model/core"
	err "driver-location-api/error"
	"net/http"
)

type DriverLocationRequest struct {
	Type     string          `json:"type" validate:"required"`
	Location core.Coordinate `json:"location" validate:"required"`
}

type SearchDriverRequest struct {
	Radius      int             `json:"radius" validate:"required"`
	Coordinates core.Coordinate `json:"coordinates" validate:"required"`
}

func (dlr DriverLocationRequest) ToDriverInfo() model.DriverInfo {
	t := dlr.Type
	longitude := dlr.Location.Longitude
	latitude := dlr.Location.Latitude

	return model.DriverInfo{
		Location: core.Location{
			Type:        t,
			Coordinates: []float64{longitude, latitude},
		},
	}
}

func (dlr DriverLocationRequest) ValidateValues() *err.Error {
	if !model.IsValidLongitude(dlr.Location.Longitude) ||
		!model.IsValidLatitude(dlr.Location.Latitude) ||
		!model.IsPointType(dlr.Type) {
		return &err.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "Make sure fields are not empty and valid",
			Details: constants.ErrorInvalidCoordinates,
		}
	}
	return nil
}
