package handler

import (
	"driver-location-api/controllers/model"
	"driver-location-api/controllers/model/dto/request"
	e "driver-location-api/error"
	"driver-location-api/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

type DriverLocationHandler interface {
	SaveDriverLocation(c *fiber.Ctx) error
	Search(c *fiber.Ctx) error
}

type driverLocationHandler struct {
	Service services.DriverLocationService
}

func NewDriverLocationHandler(service services.DriverLocationService) DriverLocationHandler {
	return driverLocationHandler{Service: service}
}

func (dlh driverLocationHandler) SaveDriverLocation(c *fiber.Ctx) error {
	var response interface{}
	ct := c.Get(fiber.HeaderContentType)

	if strings.Split(ct, ";")[0] == fiber.MIMEMultipartForm {
		fh, er := c.FormFile("drivers")
		if er != nil {
			return c.Status(http.StatusBadRequest).JSON(model.ApiResponse{
				Code:    http.StatusBadRequest,
				Message: er.Error(),
				Data:    nil,
			})
		}
		err := dlh.Service.SaveDriverLocationFile(fh)
		if err != nil {
			return c.Status(err.Code).JSON(model.ApiResponse{
				Code:    err.Code,
				Message: err.Message,
				Data:    nil,
			})
		}
		response = "uploading data"

	} else {
		var dlr request.DriverLocationRequest
		var err *e.Error

		er := c.BodyParser(&dlr)
		if er != nil {
			return c.Status(http.StatusBadRequest).JSON(er)
		}
		response, err = dlh.Service.SaveDriverLocation(dlr)

		if err != nil {
			return c.Status(err.Code).JSON(model.ApiResponse{
				Code:    err.Code,
				Message: err.Message,
				Data:    response,
			})
		}
	}

	return c.Status(http.StatusCreated).JSON(model.ApiResponse{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Data:    response,
	})

}

func (dlh driverLocationHandler) Search(c *fiber.Ctx) error {
	var sd request.SearchDriverRequest

	if err := c.BodyParser(&sd); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	response, err := dlh.Service.GetNearestDriver(sd)
	if err != nil {
		return c.Status(err.Code).JSON(model.ApiResponse{
			Code:    err.Code,
			Message: err.Message,
			Data:    response,
		})
	}

	return c.Status(http.StatusOK).JSON(model.ApiResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    response,
	})
}
