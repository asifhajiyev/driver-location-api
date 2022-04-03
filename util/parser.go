package util

import (
	"encoding/csv"
	log "github.com/sirupsen/logrus"
	"math"
	"mime/multipart"
	"strconv"
)

func StringToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Errorf("'%s' could not be parsed", s)
	}
	return f
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Errorf("'%s' could not be parsed", s)
	}
	return i
}

func CsvToSlice(fh *multipart.FileHeader) [][]string {
	f, err := fh.Open()
	if err != nil {
		log.Errorf("unable to read file %v", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Errorf("unable to parse CSV file %v", err)
	}
	return records[1:]
}

func DegreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}

func FloatToTwoDecimalFloat(f float64) float64 {
	return math.Round(f*100) / 100
}
