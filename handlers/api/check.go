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
	Statuses    []struct {
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
			if existingService.Statuses[0].Date.Before(service.Statuses[0].Date) {
				serviceMap[service.Title] = service
			}
		} else {
			serviceMap[service.Title] = service
		}
	}

	for _, service := range serviceMap {
		sort.Slice(service.Statuses, func(i, j int) bool {
			return service.Statuses[i].Date.After(service.Statuses[j].Date)
		})
	}

	var services []Check

	for _, service := range serviceMap {
		services = append(services, service)
	}

	if len(services) == 0 {
		return c.JSON([]Check{})
	}

	sort.Slice(services, func(i, j int) bool {
		return services[i].Statuses[0].Date.After(services[j].Statuses[0].Date)
	})

	for i, j := 0, len(services)-1; i < j; i, j = i+1, j-1 {
		services[i], services[j] = services[j], services[i]
	}

	return c.JSON(services)
}
