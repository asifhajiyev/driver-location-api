package handler

import (
	"driver-location-api/controllers/model"
	"driver-location-api/controllers/model/dto/request"
	e "driver-location-api/error"
	"driver-location-api/logger"
	"driver-location-api/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type DriverHandler interface {
	SaveDriverLocation(c *fiber.Ctx) error
	UploadDriverLocationFile(c *fiber.Ctx) error
	SearchDriver(c *fiber.Ctx) error
}

type driverHandler struct {
	service services.DriverService
}

func NewDriverHandler(service services.DriverService) DriverHandler {
	return driverHandler{service: service}
}

// SaveDriverLocation godoc
// @Summary			Save Driver Location
// @Tags 			Driver
// @Description 	Save Driver Location
// @Accept      	json
// @Produce     	json
// @Param       	driverLocation  body 		request.DriverLocationRequest 	false 	"driverLocation"
// @Success     	200  {object}  model.RestResponse
// @Router      	/drivers/save [post]
func (dh driverHandler) SaveDriverLocation(c *fiber.Ctx) error {
	logger.Info("SaveDriverInfo.begin")
	var dlr request.DriverLocationRequest
	var err *e.Error

	if er := c.BodyParser(&dlr); er != nil {
		logger.Error("SaveDriverInfo.error", er)
		return c.Status(http.StatusBadRequest).JSON(
			model.BuildRestResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest),
				nil, er.Error()))
	}

	vrErr := model.ValidateStructFields(dlr)
	if vrErr != nil {
		logger.Error("SaveDriverInfo.error", vrErr)
		return c.Status(http.StatusBadRequest).JSON(
			model.BuildRestResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, vrErr))
	}

	response, err := dh.service.SaveDriverLocation(dlr)
	if err != nil {
		logger.Error("SaveDriverInfo.error", err)
		return c.Status(err.Code).JSON(
			model.BuildRestResponse(err.Code, err.Message, response, err.Details))
	}

	logger.Info("SaveDriverInfo.end", response)
	return c.Status(http.StatusCreated).JSON(
		model.BuildRestResponse(http.StatusCreated, http.StatusText(http.StatusCreated), response, nil))
}

// UploadDriverLocationFile godoc
// @Summary			Upload Driver Location CSV file
// @Tags 			Driver
// @Description 	Upload Driver Location CSV file
// @Accept			multipart/form-data
// @Produce     	json
// @Param       	drivers 		formData 	file 							false 	"drivers"
// @Success     	200  {object}  model.RestResponse
// @Router      	/drivers/upload-driver-file [post]
func (dh driverHandler) UploadDriverLocationFile(c *fiber.Ctx) error {
	logger.Info("UploadDriverLocationFile.begin")
	fh, er := c.FormFile("drivers")

	if er != nil {
		logger.Error("UploadDriverLocationFile.error", er)
		return c.Status(http.StatusUnprocessableEntity).JSON(
			model.BuildRestResponse(http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity),
				nil, er.Error()))
	}

	response, err := dh.service.SaveDriverLocationFile(fh)
	if err != nil {
		logger.Error("UploadDriverLocationFile.error", err)
		return c.Status(err.Code).JSON(
			model.BuildRestResponse(err.Code, err.Message, response, nil))
	}

	logger.Info("SaveDriverInfo.end", response)
	return c.Status(http.StatusCreated).JSON(
		model.BuildRestResponse(http.StatusCreated, http.StatusText(http.StatusCreated), response, nil))
}

// SearchDriver godoc
// @Summary			Search Driver by giving rider location and maximum distance
// @Tags 			Driver
// @Description 	Search Driver
// @Accept      	json
// @Produce     	json
// @Param       	riderLocation  body 	request.SearchDriverRequest 	true 	"riderLocation and radius"
// @Success     	200  {object}  model.RestResponse
// @Router      	/drivers/search [post]
func (dh driverHandler) SearchDriver(c *fiber.Ctx) error {
	logger.Info("SearchDriver.begin")
	var sdr request.SearchDriverRequest
	if err := c.BodyParser(&sdr); err != nil {
		logger.Error("SearchDriver.error", err)
		return c.Status(http.StatusBadRequest).JSON(
			model.BuildRestResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest),
				nil, err.Error()))
	}

	if vrErr := model.ValidateStructFields(sdr); vrErr != nil {
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
