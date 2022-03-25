package dto

import (
	err "driver-location-api/error"
	"driver-location-api/model/core"
	"driver-location-api/util"
	"net/http"
)

type DriverLocationRequest struct {
	Longitude string
	Latitude  string
}

func (dlr DriverLocationRequest) ToRepoModel() core.DriverLocation {
	longitude := util.StringToFloat(dlr.Longitude)
	latitude := util.StringToFloat(dlr.Latitude)

	return core.DriverLocation{
		Coordinates: [2]float64{longitude, latitude},
	}
}

func (dlr DriverLocationRequest) Validate() *err.Error {
	if dlr.Longitude == "" || dlr.Latitude == "" ||
		util.StringToFloat(dlr.Longitude) < -180 || util.StringToFloat(dlr.Longitude) > 180 ||
		util.StringToFloat(dlr.Latitude) < -90 || util.StringToFloat(dlr.Latitude) > 90 {
		return &err.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "Make sure fields are not empty and in valid range",
		}
	}
	return nil
}
