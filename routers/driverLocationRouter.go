package routers

import (
	"github.com/gofiber/fiber/v2"
)

func (h HandlerList) SetupDriverRoute(r fiber.Router) {
	r.Post("/save", h.Dh.SaveDriverLocation)
	r.Post("/search", h.Dh.SearchDriver)
}
