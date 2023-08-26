package main

import (
	"fmt"
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"

	"tritan.dev/status-page/config"
	"tritan.dev/status-page/database"
	"tritan.dev/status-page/router"
	Sentry "tritan.dev/status-page/sentry"
	"tritan.dev/status-page/timers"
)

func main() {
	defer handlePanic()
	app := fiber.New()
	cfg := config.LoadConfig()

	db, err := database.Init()
	if err != nil {
		handleError(err)
	}

	Sentry.Init(&cfg)

	go func() {
		err := timers.Init(cfg, db)
		if err != nil {
			handleError(err)
		}
	}()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	app.Use(Sentry.Middleware())

	router.SetupRoutes(app)

	err = app.Listen(cfg.Port)
	if err != nil {
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
