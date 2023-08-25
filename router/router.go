package router

import (
	"github.com/gofiber/fiber/v2"

	"tritan.dev/status-page/handlers"
)

func SetupRoutes(app *fiber.App) fiber.Handler {
	app.Use("/api/check", handlers.Check)

	return nil
}
