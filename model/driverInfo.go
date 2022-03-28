package model

import "driver-location-api/model/core"

type DriverInfo struct {
	Location core.Location `json:"location" bson:"location"`
}
