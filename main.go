package main

import (
	"context"
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

// @title           Driver Location API
// @version         1.0
// @description     This is a Driver Location API to save them and search

// @contact.email  	asif.hajiyev@outlook.com

// @BasePath  /api/

func main() {
	InitEnvVariables()

	mongo, err := db.NewMongoRepository(os.Getenv("bitaksi_task_DB_NAME"), os.Getenv("bitaksi_task_DB_TIMEOUT"))
	if err != nil {
		log.Panic("could not connect to database")
	}

	dRepo := repositories.NewDriverRepository(mongo)
	defer db.CloseConnection(mongo.Client, context.Background())

	dService := services.NewDriverService(dRepo)
	dHandler := handler.NewDriverHandler(dService)

	app := fiber.New()

	r := routers.HandlerList{Dh: dHandler}
	r.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))

}
