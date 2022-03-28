package dto

import (
	err "driver-location-api/error"
	"driver-location-api/model"
	"driver-location-api/model/core"
	"net/http"
)

type DriverLocationRequest struct {
	Type     string          `json:"type"`
	Location core.Coordinate `json:"location"`
}

func (dlr DriverLocationRequest) ToRepoModel() core.DriverInfo {
	t := dlr.Type
	longitude := dlr.Location.Longitude
	latitude := dlr.Location.Latitude

	return core.DriverInfo{
		Location: core.Location{
			Type:        t,
			Coordinates: []float64{longitude, latitude},
		},
	}
}

func (dlr DriverLocationRequest) Validate() *err.Error {
	if !model.IsValidLongitude(dlr.Location.Longitude) ||
		!model.IsValidLatitude(dlr.Location.Latitude) ||
		!model.IsPointType(dlr.Type) {
		return &err.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "Make sure fields are not empty and valid",
		}
	}
	return nil
}
