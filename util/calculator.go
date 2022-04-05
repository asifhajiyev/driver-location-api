package util

import (
	"driver-location-api/domain/model/core"
	"math"
)

const EarthRadius = 6371

type DistanceCalculator interface {
	Calculate(from core.Coordinate, to core.Coordinate) float64
}

type CalculateByHaversine struct {
}

func (h CalculateByHaversine) Calculate(from core.Coordinate, to core.Coordinate) float64 {
	fromLng := DegreesToRadians(*from.Longitude)
	fromLt := DegreesToRadians(*from.Latitude)
	toLng := DegreesToRadians(*to.Longitude)
	toLt := DegreesToRadians(*to.Latitude)

	deltaLng := toLng - fromLng
	deltaLt := toLt - fromLt

	a := math.Pow(math.Sin(deltaLt/2), 2) + math.Cos(fromLt)*math.Cos(toLt)*
		math.Pow(math.Sin(deltaLng/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return FloatToTwoDecimalFloat(c * EarthRadius)
}
