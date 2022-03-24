package handler

import (
	"driver-location-api/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type DriverLocationHandler interface {
	SaveDriverLocation(c *fiber.Ctx) error
}

type DriverLocationHandlerImpl struct {
	dlc service.DriverLocationService
}

func NewDriverLocationHandler(s service.DriverLocationService) DriverLocationHandler {
	return &DriverLocationHandlerImpl{dlc: s}
}

func (dlh DriverLocationHandlerImpl) SaveDriverLocation(c *fiber.Ctx) error {
	fmt.Println("it works")
	return nil
}
