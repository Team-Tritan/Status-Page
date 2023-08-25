package handlers

import (
	"fmt"
	"net"
	"time"

	"github.com/gofiber/fiber/v2"
	"tritan.dev/status-page/config"
)

type ServiceStatus struct {
	Title       string
	Hostname    string
	Port        string
	Description string
	Status      string
	Latency     string
}

func Check(c *fiber.Ctx) error {
	cfg := config.LoadConfig()

	var services []ServiceStatus

	for _, service := range cfg.Services.Services {
		address := fmt.Sprintf("%s:%s", service.Hostname, service.Port)
		start := time.Now()
		_, err := net.DialTimeout("tcp", address, 2*time.Second)
		elapsed := time.Since(start).Milliseconds()

		status := "OK"
		if err != nil {
			status = "FAIL"
		}

		services = append(services, ServiceStatus{
			Title:       service.Title,
			Hostname:    service.Hostname,
			Port:        service.Port,
			Description: service.Description,
			Status:      status,
			Latency:     fmt.Sprintf("%dms", elapsed),
		})
	}

	return c.JSON(services)
}
