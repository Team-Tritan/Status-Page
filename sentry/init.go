package sentry

import (
	"log"

	"github.com/getsentry/sentry-go"

	"tritan.dev/status-page/config"
)

func Init(config *config.Config) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              config.Sentry,
		TracesSampleRate: 1.0,
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}
