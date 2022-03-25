package handler

import (
	"driver-location-api/model/dto"
	"driver-location-api/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type DriverLocationHandler interface {
	SaveDriverLocation(c *fiber.Ctx) error
}

type driverLocationHandler struct {
	Dls service.DriverLocationService
}

func NewDriverLocationHandler(Dls service.DriverLocationService) DriverLocationHandler {
	return driverLocationHandler{Dls: Dls}
}

func (dlh driverLocationHandler) SaveDriverLocation(c *fiber.Ctx) error {
	var dlr dto.DriverLocationRequest

	if err := c.BodyParser(&dlr); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	response, err := dlh.Dls.SaveDriverLocation(dlr)
	if err != nil {
		return c.Status(err.Code).JSON(err)
	}
	return c.Status(http.StatusCreated).JSON(ApiResponse{
		Code:    http.StatusCreated,
		Message: "success",
		Data:    response,
	})
}
