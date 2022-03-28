package model

type RideInfo struct {
	DriverInfo DriverInfo `json:"driverInfo"`
	Distance   float64    `json:"distance"`
}
