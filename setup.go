package main

import (
	"driver-location-api/model/core"
	"driver-location-api/repository"
	"driver-location-api/util"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func InitEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln("could not load env", err)
	}
}

type Uploader interface {
	UploadDriverLocationFile(di []core.DriverInfo)
}

type upload struct {
	repo repository.DriverLocationRepo
}

func NewUpload(repo repository.DriverLocationRepo) Uploader {
	return upload{repo: repo}
}

func (u upload) UploadDriverLocationFile(di []core.DriverInfo) {
	u.repo.SaveDriverLocationFile(di)
}

func addToSliceAndUpload(dlRepo repository.DriverLocationRepo, s [][]string) {
	var dls []core.DriverInfo

	for i := 0; i < len(s); i++ {
		longitude := util.StringToFloat(s[i][0])
		latitude := util.StringToFloat(s[i][1])
		di := core.DriverInfo{Location: core.Location{
			Type:        "Point",
			Coordinates: []float64{longitude, latitude},
		}}
		dls = append(dls, di)
	}

	u := NewUpload(dlRepo)
	u.UploadDriverLocationFile(dls)
}

func UploadDriverData(dlRepo repository.DriverLocationRepo, fp string) {
	var dlUploadPatchSize = util.StringToInt(os.Getenv("bitaksi_task_INSERT_DOC_NUM_AT_ONCE"))
	data := util.CsvToSlice(fp)
	patchData := make([][]string, 0)

	for _, v := range data {
		patchData = append(patchData, v)
		if len(patchData) == dlUploadPatchSize {
			fmt.Println("data in len", len(patchData))
			go addToSliceAndUpload(dlRepo, patchData)
			patchData = nil
		}
	}
	if len(patchData) > 0 {
		fmt.Println("data in remainder", len(patchData))
		go addToSliceAndUpload(dlRepo, patchData)
	}
}

/*func parseCsvAndUploadPatch(dlRepo repository.DriverLocationRepo, file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	data := make([][]string, 0)
	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			go addToSliceAndUpload(dlRepo, data)
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, record)

		if len(data) == 50000 {
			fmt.Println("data in len", len(data))
			go addToSliceAndUpload(dlRepo, data)
			data = nil
		}
	}
}*/
