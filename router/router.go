package router

import (
	"github.com/gofiber/fiber/v2"

	api "tritan.dev/status-page/handlers/api"
)

func SetupRoutes(app *fiber.App) fiber.Handler {

	app.Use("/api/check", api.CheckEndpoint)

	return nil
}
