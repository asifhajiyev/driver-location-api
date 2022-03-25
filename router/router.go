package router

import (
	"driver-location-api/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type HandlerList struct {
	Dlh handler.DriverLocationHandler
}

func (h *HandlerList) SetupRoutes(app *fiber.App) {
	app.Use(logger.New())

	dl := app.Group("api").Group("driver-location")
	h.SetupDriverLocationRoute(dl)
}
