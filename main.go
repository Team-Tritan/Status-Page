package main

import (
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"

	"tritan.dev/status-page/config"
	"tritan.dev/status-page/router"
	Sentry "tritan.dev/status-page/sentry"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Panic: %v", err)
			sentry.CaptureException(fmt.Errorf("%v", err))
			sentry.Flush(2 * time.Second)
		}
	}()

	app := fiber.New()
	config := config.LoadConfig()

	Sentry.Init(&config)
	app.Use(Sentry.Middleware())
	app.Use(router.SetupRoutes(app))

	err := app.Listen(config.Port)
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}

}
