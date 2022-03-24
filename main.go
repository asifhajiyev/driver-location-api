package main

import (
	"driver-location-api/db"
	"driver-location-api/handler"
	"driver-location-api/repository"
	"driver-location-api/router"
	"driver-location-api/service"
	"driver-location-api/util"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	util.InitEnvVariables()

	mongo, err := db.NewMongoRepository(os.Getenv("bitaksi_task_DB_NAME"), os.Getenv("bitaksi_task_DB_TIMEOUT"))
	if err != nil {
		log.Panic("could not connect to database")
	}

	dlRepo := repository.NewDriverLocationRepo(mongo)
	dlService := service.NewDriverLocationServiceImpl(dlRepo)
	dlHandler := handler.NewDriverLocationHandler(dlService)

	app := fiber.New()
	r := router.HandlerList{Dlh: dlHandler}
	r.SetupRoutes(app)

	/*c := db.GetCollection("driver_locations")

	driver := core.DriverLocation{Coordinates: [2]float64{23.20, 22.10}}

	_, err := c.InsertOne(context.Background(), driver)
	if err != nil {
		fmt.Println("in insert: ", err)
	}*/

	/*app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "it works"})
	})*/

	log.Fatal(app.Listen(":8080"))
}
