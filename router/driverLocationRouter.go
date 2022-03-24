package router

import (
	"driver-location-api/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupDriverLocationRoute(dl fiber.Router) {
	dl.Post("/save", handler.DriverLocationHandlerImpl{}.SaveDriverLocation)
}
