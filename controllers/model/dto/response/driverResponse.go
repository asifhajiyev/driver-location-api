package response

import (
	"driver-location-api/domain/model/core"
)

type DriverLocationResponse struct {
	Type     string          `json:"type"`
	Location core.Coordinate `json:"location"`
}
