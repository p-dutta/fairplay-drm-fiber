package router

import (
	"encoding/json"
	"fairplay-ksm/handler"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {

	api := app.Group(os.Getenv("ROUTE_PREFIX") + "/" + os.Getenv("API_VERSION"))
	// api := app.Group("/" + os.Getenv("API_VERSION"))

	// FPS
	// fps := api.Group("/fairplay")
	api.Get("/health", HealthCheckHandler)
	api.Post("/license", handler.GetLicense)

}

func HealthCheckHandler(c *fiber.Ctx) error {
	// Create a map representing the JSON response
	response := map[string]string{
		"status":  "ok",
		"message": "Service is healthy",
	}

	// Marshal the map into JSON format
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Internal server error")
	}

	// Set the Content-Type header and write the JSON response
	c.Set("Content-Type", "application/json")
	return c.Status(http.StatusOK).Send(jsonResponse)
}
