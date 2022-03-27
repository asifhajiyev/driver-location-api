package dto

import (
	err "driver-location-api/error"
	"driver-location-api/model"
	"driver-location-api/model/core"
	"driver-location-api/util"
	"net/http"
)

type DriverLocationRequest struct {
	Type      string `json:"type"`
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

func (dlr DriverLocationRequest) ToRepoModel() core.DriverLocation {
	t := dlr.Type
	longitude := util.StringToFloat(dlr.Longitude)
	latitude := util.StringToFloat(dlr.Latitude)

	return core.DriverLocation{
		Location: core.Geometry{
			Type:        t,
			Coordinates: []float64{longitude, latitude},
		},
	}
}

func (dlr DriverLocationRequest) Validate() *err.Error {
	if !model.IsValidLongitude(util.StringToFloat(dlr.Longitude)) ||
		!model.IsValidLatitude(util.StringToFloat(dlr.Latitude)) ||
		!model.IsPointType(dlr.Type) {
		return &err.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "Make sure fields are not empty and valid",
		}
	}
	return nil
}
