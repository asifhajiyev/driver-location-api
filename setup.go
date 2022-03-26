package main

import (
	"driver-location-api/model/core"
	"driver-location-api/repository"
	"github.com/joho/godotenv"
	"log"
)

func InitEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln("could not load env", err)
	}
}

type Uploader interface {
	UploadDriverLocationFile(dl []core.DriverLocation)
}

type upload struct {
	repo repository.DriverLocationRepo
}

func NewUpload(repo repository.DriverLocationRepo) Uploader {
	return update{repo: repo}
}
