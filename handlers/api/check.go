package handlers

import (
	"context"
	"sort"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	db := c.Locals("db").(*mongo.Client)

	serviceMap, err := getServiceMap(db)
	if err != nil {
		return err
	}

	services := getSortedServices(serviceMap)

	return c.JSON(services)
}

func getServiceMap(db *mongo.Client) (map[string]Check, error) {
	serviceMap := make(map[string]Check)

	collection := db.Database("status-page").Collection("services")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return serviceMap, nil
		}
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var service Check
		if err := cursor.Decode(&service); err != nil {
			return nil, err
		}

		if existingService, ok := serviceMap[service.Title]; ok {
			if existingService.Statuses[0].Date.Before(service.Statuses[0].Date) {
				serviceMap[service.Title] = service
			}
		} else {
			serviceMap[service.Title] = service
		}
	}

	return serviceMap, nil
}

func getSortedServices(serviceMap map[string]Check) []Check {
	var services []Check

	for _, service := range serviceMap {
		sort.Slice(service.Statuses, func(i, j int) bool {
			return service.Statuses[i].Date.After(service.Statuses[j].Date)
		})
		services = append(services, service)
	}

	if len(services) == 0 {
		return []Check{}
	}

	sort.Slice(services, func(i, j int) bool {
		return services[i].Statuses[0].Date.After(services[j].Statuses[0].Date)
	})

	for i, j := 0, len(services)-1; i < j; i, j = i+1, j-1 {
		services[i], services[j] = services[j], services[i]
	}

	return services
}
