package main

import (
	"driver-location-api/db"
	"driver-location-api/handler"
	"driver-location-api/model/core"
	"driver-location-api/repository"
	"driver-location-api/router"
	"driver-location-api/service"
	"driver-location-api/util"
	"encoding/csv"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"os"
)

func main() {
	InitEnvVariables()

	mongo, err := db.NewMongoRepository(os.Getenv("bitaksi_task_DB_NAME"), os.Getenv("bitaksi_task_DB_TIMEOUT"))
	if err != nil {
		log.Panic("could not connect to database")
	}

	dlRepo := repository.NewDriverLocationRepo(mongo)
	dlService := service.NewDriverLocationService(dlRepo)
	dlHandler := handler.NewDriverLocationHandler(dlService)

	UploadDriverData(dlRepo)

	app := fiber.New()
	r := router.HandlerList{Dlh: dlHandler}
	r.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))

}

func insert(dlRepo repository.DriverLocationRepo, s [][]string, c chan int) {
	var dls []core.DriverLocation
	for i := 0; i < len(s); i++ {
		longitude := util.StringToFloat(s[i][0])
		latitude := util.StringToFloat(s[i][1])
		dl := core.DriverLocation{Location: core.Geometry{
			Type:        "Point",
			Coordinates: []float64{longitude, latitude},
		}}
		dls = append(dls, dl)
	}

	//fmt.Println("dls is", dls)

	u := NewUpdate(dlRepo)
	u.UploadDriverLocationFile(dls)

	c <- len(s)
}

/*func insert(dlRepo repository.DriverLocationRepo, s [][]string) {
	var dls []core.DriverLocation
	//fmt.Println("s is", s)
	for i := 0; i < len(s); i++ {
		longitude := util.StringToFloat(s[i][0])
		latitude := util.StringToFloat(s[i][1])
		dl := core.DriverLocation{Location: core.Geometry{
			Type:        "Point",
			Coordinates: []float64{longitude, latitude},
		}}
		dls = append(dls, dl)
	}

	//fmt.Println("dls is", dls)

	u := NewUpdate(dlRepo)
	u.UploadDriverLocationFile(dls)
}*/

func UploadDriverData(dlRepo repository.DriverLocationRepo) {
	fp := "C:\\Users\\HaciyevAB\\Downloads\\driver_locations_test\\Coordinates.csv"
	parse(dlRepo, fp)
}

func parse(dlRepo repository.DriverLocationRepo, file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	data := make([][]string, 0)
	c := make(chan int)
	r := csv.NewReader(f)
	counter := 0

	for {
		record, err := r.Read()

		if err == io.EOF {
			fmt.Println("data in eof", len(data))
			go insert(dlRepo, data, c)
			//insert(dlRepo, data)
			counter++
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		data = append(data, record)

		if len(data) == 50000 {
			fmt.Println("data in len", len(data))
			go insert(dlRepo, data, c)
			counter++
			data = nil
		}
	}
}

type Update interface {
	UploadDriverLocationFile(dl []core.DriverLocation)
}

type update struct {
	repo repository.DriverLocationRepo
}

func NewUpdate(repo repository.DriverLocationRepo) Update {
	return update{repo: repo}
}

func (u update) UploadDriverLocationFile(dl []core.DriverLocation) {
	u.repo.SaveDriverLocationFile(dl)
}
