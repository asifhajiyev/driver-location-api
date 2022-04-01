package main

import (
	"driver-location-api/controllers/handler"
	"driver-location-api/db"
	"driver-location-api/repositories"
	"driver-location-api/router"
	"driver-location-api/service"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	InitEnvVariables()

	mongo, err := db.NewMongoRepository(os.Getenv("bitaksi_task_DB_NAME"), os.Getenv("bitaksi_task_DB_TIMEOUT"))
	if err != nil {
		log.Panic("could not connect to database")
	}

	dlRepo := repositories.NewDriverLocationRepo(mongo)
	dlService := service.NewDriverLocationService(dlRepo)
	dlHandler := handler.NewDriverLocationHandler(dlService)

	app := fiber.New()

	/*app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))*/

	r := router.HandlerList{Dlh: dlHandler}
	r.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))

}
