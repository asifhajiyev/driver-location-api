package core

import "fmt"

type Location struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

type Coordinate struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func NewPoint(longitude, latitude float64) Location {
	return Location{
		"Point",
		[]float64{longitude, latitude},
	}
}

func GetCoordinates(l Location) Coordinate {
	fmt.Println("l is", l.Coordinates)
	return Coordinate{
		Longitude: l.Coordinates[0],
		Latitude:  l.Coordinates[1],
	}
}

/*func GetCoordinates(l Location) Coordinate {
	fmt.Println("l is", l.Coordinates)
	return Coordinate{
		Longitude: l.Coordinates.([]float64)[0],
		Latitude:  l.Coordinates.([]float64)[1],
	}
}*/
