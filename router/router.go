package router

import (
	"github.com/gofiber/fiber/v2"

	api "tritan.dev/status-page/handlers/api"
	ui "tritan.dev/status-page/handlers/ui"
)

func SetupRoutes(app *fiber.App) fiber.Handler {
	app.Get("/", ui.DisplayUI)
	app.Get("/api/check", api.CheckEndpoint)

	return nil
}
