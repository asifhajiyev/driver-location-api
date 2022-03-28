package model

import "driver-location-api/model/core"

type RideInfo struct {
	DriverInfo core.DriverInfo `json:"driverInfo"`
	Distance   float64         `json:"distance"`
}
