package handlers

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"tritan.dev/status-page/config"
)

type Check struct {
	Title       string
	Hostname    string
	Port        string
	Description string
	Checks      []struct {
		Status  string
		Latency string
		Date    time.Time
	}
}

func CheckEndpoint(c *fiber.Ctx) error {
	cfg := config.LoadConfig()

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		return err
	}

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			fmt.Println(err)
		}
	}()

	serviceMap := make(map[string]Check)

	collection := client.Database("status-page").Collection("services")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   "true",
				"message": "No data found",
			})
		}
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var service Check
		if err := cursor.Decode(&service); err != nil {
			return err
		}

		if existingService, ok := serviceMap[service.Title]; ok {
			if existingService.Checks[0].Date.Before(service.Checks[0].Date) {
				serviceMap[service.Title] = service
			}
		} else {
			serviceMap[service.Title] = service
		}
	}

	for _, service := range serviceMap {
		sort.Slice(service.Checks, func(i, j int) bool {
			return service.Checks[i].Date.After(service.Checks[j].Date)
		})
	}

	var services []Check

	for _, service := range serviceMap {
		services = append(services, service)
	}

	sort.Slice(services, func(i, j int) bool {
		return services[i].Checks[0].Date.After(services[j].Checks[0].Date)
	})

	return c.JSON(services)
}
