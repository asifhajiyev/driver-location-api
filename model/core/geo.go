package core

type Geometry struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}

func NewPoint(longitude, latitude float64) Geometry {
	return Geometry{
		"Point",
		[]float64{longitude, latitude},
	}
}
