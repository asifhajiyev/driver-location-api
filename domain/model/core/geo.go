package core

import (
	"driver-location-api/domain/constants"
)

type Location struct {
	Type        string     `json:"type" bson:"type"`
	Coordinates []*float64 `json:"coordinates" bson:"coordinates"`
}

type Coordinate struct {
	Longitude *float64 `json:"longitude" validate:"required"`
	Latitude  *float64 `json:"latitude" validate:"required"`
}

func NewPoint(longitude, latitude *float64) Location {
	return Location{
		constants.LocationTypePoint,
		[]*float64{longitude, latitude},
	}
}

func GetCoordinates(l Location) Coordinate {
	return Coordinate{
		Longitude: l.Coordinates[0],
		Latitude:  l.Coordinates[1],
	}
}
