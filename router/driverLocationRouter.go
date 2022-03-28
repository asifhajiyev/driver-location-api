package router

import (
	"github.com/gofiber/fiber/v2"
)

func (h HandlerList) SetupDriverLocationRoute(dl fiber.Router) {
	dl.Post("/save", h.Dlh.SaveDriverLocation)
	dl.Post("/search", h.Dlh.Search)
}
