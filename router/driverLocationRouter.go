package router

import (
	"github.com/gofiber/fiber/v2"
)

func (h HandlerList) SetupDriverRoute(dl fiber.Router) {
	dl.Post("/save", h.Dlh.SaveDriverLocation)
	dl.Post("/search", h.Dlh.Search)
}
