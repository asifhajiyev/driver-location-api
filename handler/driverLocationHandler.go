package handler

import (
	"driver-location-api/model/dto"
	"driver-location-api/service"
	"driver-location-api/util"
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

func (dlh driverLocationHandler) Search(c *fiber.Ctx) error {
	longitude := util.StringToFloat(c.Query("longitude"))
	latitude := util.StringToFloat(c.Query("latitude"))
	radius := util.StringToInt(c.Query("radius"))
	response, err := dlh.Dls.GetNearestDriver(longitude, latitude, radius)
	if err != nil {
		return c.Status(err.Code).JSON(err)
	}
	return c.Status(http.StatusOK).JSON(ApiResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    response,
	})
}
