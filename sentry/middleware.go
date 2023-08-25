package sentry

import (
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
)

func Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			sentry.CaptureException(err)
			sentry.Flush(2 * time.Second)
		}
		return err
	}
}
