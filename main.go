package main

import (
	"driver-location-api/controllers/handler"
	"driver-location-api/db"
	_ "driver-location-api/docs"
	"driver-location-api/repositories"
	"driver-location-api/routers"
	"driver-location-api/services"
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
	dlService := services.NewDriverLocationService(dlRepo)
	dlHandler := handler.NewDriverLocationHandler(dlService)

	app := fiber.New()

	r := routers.HandlerList{Dlh: dlHandler}
	r.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))

}
