package router

import (
	"github.com/gofiber/fiber/v2"
)

func (h HandlerList) SetupDriverLocationRoute(dl fiber.Router) {
	dl.Post("/save", h.Dlh.SaveDriverLocation)
	dl.Get("/search", h.Dlh.Search)
}
