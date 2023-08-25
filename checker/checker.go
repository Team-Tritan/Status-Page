package checker

import (
	"context"
	"fmt"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
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

func Init(cfg config.Config, client *mongo.Client) {
	for {
		var checks []Check // Create an array of Check objects

		for _, service := range cfg.Services.Services {
			address := fmt.Sprintf("%s:%s", service.Hostname, service.Port)
			start := time.Now()
			_, err := net.DialTimeout("tcp", address, 2*time.Second)
			elapsed := time.Since(start).Milliseconds()

			status := "OK"
			if err != nil {
				status = "FAIL"
			}

			check := Check{
				Title:       service.Title,
				Hostname:    service.Hostname,
				Port:        service.Port,
				Description: service.Description,
				Checks: []struct {
					Status  string
					Latency string
					Date    time.Time
				}{
					{
						Status:  status,
						Latency: fmt.Sprintf("%dms", elapsed),
						Date:    time.Now(),
					},
				},
			}

			checks = append(checks, check)
		}

		collection := client.Database("status-page").Collection("services")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		documents := make([]interface{}, len(checks))
		for i, check := range checks {
			documents[i] = check
		}

		_, err := collection.InsertMany(ctx, documents)
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(1 * time.Minute)
	}
}

func StartChecker(cfg config.Config, client *mongo.Client) {
	go Init(cfg, client)
}
