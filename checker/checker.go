package checker

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func Init(cfg config.Config, client *mongo.Client) {
	var mu sync.Mutex

	for {
		for _, service := range cfg.Services.Services {
			address := fmt.Sprintf("%s:%s", service.Hostname, service.Port)
			start := time.Now()
			_, err := net.DialTimeout("tcp", address, 2*time.Second)
			elapsed := time.Since(start).Milliseconds()

			status := "OK"
			if err != nil {
				status = "FAIL"
			}

			check := &Check{
				Title:       service.Title,
				Hostname:    service.Hostname,
				Port:        service.Port,
				Description: service.Description,
			}

			mu.Lock()

			collection := client.Database("status-page").Collection("services")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			var existingService Check
			err = collection.FindOne(ctx, bson.M{"title": check.Title}).Decode(&existingService)

			if err == nil {
				existingService.Statuses = append(existingService.Statuses, struct {
					Status  string
					Latency string
					Date    time.Time
				}{
					Status:  status,
					Latency: fmt.Sprintf("%dms", elapsed),
					Date:    time.Now(),
				})

				_, err := collection.UpdateOne(
					ctx,
					bson.M{"title": existingService.Title},
					bson.M{"$set": bson.M{"statuses": existingService.Statuses}},
				)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				check.Statuses = []struct {
					Status  string
					Latency string
					Date    time.Time
				}{{
					Status:  status,
					Latency: fmt.Sprintf("%dms", elapsed),
					Date:    time.Now(),
				}}

				_, err := collection.InsertOne(ctx, check)
				if err != nil {
					fmt.Println(err)
				}
			}

			mu.Unlock()
		}

		time.Sleep(1 * time.Minute)
	}
}
