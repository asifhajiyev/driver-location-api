package model

import "driver-location-api/model/core"

type DriverInfo struct {
	DriverLocation core.DriverLocation `json:"driverLocation"`
	Distance       float64             `json:"distance"`
}
