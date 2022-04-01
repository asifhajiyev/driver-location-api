package model

import (
	"driver-location-api/domain/model/core"
)

type DriverInfo struct {
	Location core.Location `json:"location" bson:"location"`
}
