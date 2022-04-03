package handler

import (
	"driver-location-api/controllers/model"
	"driver-location-api/controllers/model/dto/request"
	e "driver-location-api/error"
	"driver-location-api/logger"
	"driver-location-api/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

type DriverHandler interface {
	SaveDriverLocation(c *fiber.Ctx) error
	SearchDriver(c *fiber.Ctx) error
}

type driverHandler struct {
	service services.DriverService
}

func NewDriverHandler(service services.DriverService) DriverHandler {
	return driverHandler{service: service}
}

// SaveDriverLocation godoc
// @Summary      Save Driver Location, supports batch upload and single location object
// @Tags 		 Driver
// @Description  Save Driver Location
// @Accept       json
// @Produce      json
// @Param        driverLocation  body request.DriverLocationRequest false "driverLocation"
// @Success      200  {object}  model.RestResponse
// @Router       /api/drivers/save [post]
func (dh driverHandler) SaveDriverLocation(c *fiber.Ctx) error {
	logger.Info("SaveDriverLocation.begin")
	var response interface{}
	var err *e.Error
	ct := c.Get(fiber.HeaderContentType)

	if strings.Split(ct, ";")[0] == fiber.MIMEMultipartForm {
		logger.Info("SaveDriverLocation.batch")
		fh, er := c.FormFile("drivers")
		if er != nil {
			logger.Error("SaveDriverLocation.error", er)
			return c.Status(http.StatusUnprocessableEntity).JSON(
				model.BuildRestResponse(http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity),
					nil, er.Error()))
		}

		if response, err = dh.service.SaveDriverLocationFile(fh); err != nil {
			logger.Error("SaveDriverLocation.error", err)
			return c.Status(err.Code).JSON(
				model.BuildRestResponse(err.Code, err.Message, response, nil))
		}

	} else {
		var dlr request.DriverLocationRequest
		var err *e.Error

		if er := c.BodyParser(&dlr); er != nil {
			logger.Error("SaveDriverLocation.error", er)
			return c.Status(http.StatusBadRequest).JSON(
				model.BuildRestResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest),
					nil, er.Error()))
		}

		if vrErr := model.ValidateRequest(dlr); vrErr != nil {
			logger.Error("SaveDriverLocation.error", vrErr)
			return c.Status(http.StatusBadRequest).JSON(
				model.BuildRestResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, vrErr))
		}

		response, err = dh.service.SaveDriverLocation(dlr)
		if err != nil {
			logger.Error("SaveDriverLocation.error", err)
			return c.Status(err.Code).JSON(
				model.BuildRestResponse(err.Code, err.Message, response, err.Details))
		}
	}
	logger.Info("SaveDriverLocation.end", response)
	return c.Status(http.StatusCreated).JSON(
		model.BuildRestResponse(http.StatusCreated, http.StatusText(http.StatusCreated), response, nil))
}

func (dh driverHandler) SearchDriver(c *fiber.Ctx) error {
	fmt.Println()
	logger.Info("SearchDriver.begin")
	var sdr request.SearchDriverRequest
	if err := c.BodyParser(&sdr); err != nil {
		logger.Error("SearchDriver.error", err)
		return c.Status(http.StatusBadRequest).JSON(
			model.BuildRestResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest),
				nil, err.Error()))
	}

	if vrErr := model.ValidateRequest(sdr); vrErr != nil {
		logger.Error("SearchDriver.error", vrErr)
		return c.Status(http.StatusBadRequest).JSON(
			model.BuildRestResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest),
				nil, vrErr))
	}

	response, err := dh.service.GetNearestDriver(sdr)
	if err != nil {
		logger.Error("SearchDriver.error", err)
		return c.Status(err.Code).JSON(
			model.BuildRestResponse(err.Code, err.Message, nil, err.Details))
	}
	logger.Info("SearchDriver.end", response)
	return c.Status(http.StatusOK).JSON(
		model.BuildRestResponse(http.StatusOK, http.StatusText(http.StatusOK), response, nil))
}
