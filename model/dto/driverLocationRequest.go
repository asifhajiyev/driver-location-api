package dto

import (
	err "driver-location-api/error"
	"driver-location-api/model/core"
	"driver-location-api/util"
	"net/http"
)

type DriverLocationRequest struct {
	longitude string
	latitude  string
}

func (dlr DriverLocationRequest) ToRepoModel() core.DriverLocation {
	longitude := util.StringToFloat(dlr.longitude)
	latitude := util.StringToFloat(dlr.latitude)

	return core.DriverLocation{
		Coordinates: [2]float64{longitude, latitude},
	}
}

func (dlr DriverLocationRequest) Validate() *err.Error {
	if dlr.longitude == "" || dlr.latitude == "" ||
		util.StringToFloat(dlr.longitude) < -180 || util.StringToFloat(dlr.longitude) > 180 ||
		util.StringToFloat(dlr.latitude) < -90 || util.StringToFloat(dlr.latitude) > 90 {
		return &err.Error{
			Code:    http.StatusUnprocessableEntity,
			Message: "Make sure fields are not empty and in valid range",
		}
	}
	return nil
}
