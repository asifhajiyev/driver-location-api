package routers

import (
	"driver-location-api/middleware"
	"github.com/gofiber/fiber/v2"
)

func (h HandlerList) SetupDriverRoute(r fiber.Router) {
	r.Post("/save", h.Dh.SaveDriverLocation)
	r.Post("/upload-driver-file", h.Dh.UploadDriverLocationFile)
	r.Post("/search", middleware.BasicAuthProtector(), h.Dh.SearchDriver)
}
