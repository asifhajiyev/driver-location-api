package util

import (
	"encoding/csv"
	log "github.com/sirupsen/logrus"
	"os"
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

func ParseCSVToSlice(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("unable to parse CSV file for "+filePath, err)
	}
	return records[1:]
}
