package routers

import (
	"driver-location-api/controllers/handler"
	"driver-location-api/controllers/model"
	err "driver-location-api/error"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type HandlerList struct {
	Dlh handler.DriverLocationHandler
}

func (h *HandlerList) SetupRoutes(app *fiber.App) {
	UseSwagger(app)
	app.Use(logger.New())

	dl := app.Group("api").Group("drivers")
	h.SetupDriverRoute(dl)

	handleNotFoundError(app)
}

func handleNotFoundError(app *fiber.App) {
	app.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(model.RestResponse{
				Code:    fiber.StatusNotFound,
				Message: err.URLNotFound,
				Data:    nil,
			})
		},
	)
}

func UseSwagger(app *fiber.App) {
	route := app.Group("/swagger")
	route.Get("*", swagger.HandlerDefault)
}
