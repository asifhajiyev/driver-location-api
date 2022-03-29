package handler

import (
	"driver-location-api/model/dto/request"
	"driver-location-api/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type DriverLocationHandler interface {
	SaveDriverLocation(c *fiber.Ctx) error
	Search(c *fiber.Ctx) error
}

type driverLocationHandler struct {
	Dls service.DriverLocationService
}

func NewDriverLocationHandler(Dls service.DriverLocationService) DriverLocationHandler {
	return driverLocationHandler{Dls: Dls}
}

func (dlh driverLocationHandler) SaveDriverLocation(c *fiber.Ctx) error {
	var dlr request.DriverLocationRequest

	if err := c.BodyParser(&dlr); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	response, err := dlh.Dls.SaveDriverLocation(dlr)
	if err != nil {
		return c.Status(err.Code).JSON(ApiResponse{
			Code:    err.Code,
			Message: err.Message,
			Data:    response,
		})
	}
	return c.Status(http.StatusCreated).JSON(ApiResponse{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Data:    response,
	})
}

func (dlh driverLocationHandler) Search(c *fiber.Ctx) error {
	var sd request.SearchDriver

	if err := c.BodyParser(&sd); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	response, err := dlh.Dls.GetNearestDriver(sd)
	if err != nil {
		return c.Status(err.Code).JSON(ApiResponse{
			Code:    err.Code,
			Message: err.Message,
			Data:    response,
		})
	}

	return c.Status(http.StatusOK).JSON(ApiResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    response,
	})
}
