package util

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
)

func StringToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Panicf("could not be parsed %s", s)
	}
	return f
}

func StringToInt(s string) int {
	fmt.Println("s is", s)
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panicf("could not be parsed %s", s)
	}
	return i
}
