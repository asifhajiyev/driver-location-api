package middleware

import (
	"driver-location-api/controllers/model"
	"driver-location-api/domain/constants"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"net/http"
	"os"
)

func BasicAuthProtector() func(*fiber.Ctx) error {
	config := basicauth.Config{
		Users: map[string]string{
			os.Getenv("bitaksi_task_matchingapi_username"): os.Getenv("bitaksi_task_matchingapi_password"),
		},
		Unauthorized: basicAuthUnauthorized,
	}
	return basicauth.New(config)
}

func basicAuthUnauthorized(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(
		model.BuildRestResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized),
			nil, constants.ErrorWrongCredentialsForClients))
}
