package dto

import "driver-location-api/model/core"

type DriverLocationResponse struct {
	Type     string          `json:"type"`
	Location core.Coordinate `json:"location"`
}
