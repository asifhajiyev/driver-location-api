package model

func IsValidLongitude(longitude float64) bool {
	return longitude >= -180 && longitude <= 180
}

func IsValidLatitude(latitude float64) bool {
	return latitude >= -90 && latitude <= 90
}

func IsPointType(s string) bool {
	return s == "Point"
}
