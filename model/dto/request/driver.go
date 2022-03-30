package request

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

type SearchDriver struct {
	Radius      int             `json:"radius"`
	Coordinates core.Coordinate `json:"coordinates"`
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
