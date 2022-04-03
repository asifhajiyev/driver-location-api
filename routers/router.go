package routers

import (
	"driver-location-api/controllers/handler"
	"driver-location-api/controllers/model"
	"driver-location-api/domain/constants"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"net/http"
)

type HandlerList struct {
	Dh handler.DriverHandler
}

func (h *HandlerList) SetupRoutes(app *fiber.App) {
	useSwagger(app)
	app.Use(logger.New())

	driverRoute := app.Group("api").Group("drivers")
	h.SetupDriverRoute(driverRoute)

	handleNotFoundError(app)
}

func handleNotFoundError(app *fiber.App) {
	app.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(
				model.BuildRestResponse(fiber.StatusNotFound, http.StatusText(fiber.StatusNotFound),
					nil, constants.ErrorURLNotFound))
		},
	)
}

func useSwagger(app *fiber.App) {
	route := app.Group("/swagger")
	route.Get("*", swagger.HandlerDefault)
}
