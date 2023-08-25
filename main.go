package main

import (
	"fmt"
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"

	"tritan.dev/status-page/checker"
	"tritan.dev/status-page/config"
	"tritan.dev/status-page/database"
	"tritan.dev/status-page/router"
	Sentry "tritan.dev/status-page/sentry"
)

func main() {
	defer handlePanic()

	app := fiber.New()
	cfg := config.LoadConfig()

	db, err := database.Connect()
	if err != nil {
		handleError(err)
	}

	Sentry.Init(&cfg)
	go checker.Init(cfg, db)

	app.Use(Sentry.Middleware())
	app.Use(router.SetupRoutes(app))

	if err := app.Listen(cfg.Port); err != nil {
		handleError(err)
	}
}

func handlePanic() {
	if err := recover(); err != nil {
		log.Printf("Panic: %v", err)
		sentry.CaptureException(fmt.Errorf("%v", err))
	}
}

func handleError(err error) {
	sentry.CaptureException(err)
	log.Fatal(err)
}
