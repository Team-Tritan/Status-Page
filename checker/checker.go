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
	Statuses    []struct {
		Status  string
		Latency string
		Date    time.Time
	}
}

func Init(cfg config.Config, client *mongo.Client) {
	for {
		var checks []*Check // Change the slice type to []*Check

		for _, service := range cfg.Services.Services {
			address := fmt.Sprintf("%s:%s", service.Hostname, service.Port)
			start := time.Now()
			_, err := net.DialTimeout("tcp", address, 2*time.Second)
			elapsed := time.Since(start).Milliseconds()

			status := "OK"
			if err != nil {
				status = "FAIL"
			}

			var check *Check
			for i := range checks {
				if checks[i].Title == service.Title {
					check = checks[i]
					break
				}
			}

			if check == nil {
				check = &Check{
					Title:       service.Title,
					Hostname:    service.Hostname,
					Port:        service.Port,
					Description: service.Description,
				}
				checks = append(checks, check)
			}

			check.Statuses = append(check.Statuses, struct {
				Status  string
				Latency string
				Date    time.Time
			}{
				Status:  status,
				Latency: fmt.Sprintf("%dms", elapsed),
				Date:    time.Now(),
			})
		}

		collection := client.Database("status-page").Collection("services")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if len(checks) > 0 {
			documents := make([]interface{}, len(checks))
			for i, check := range checks {
				documents[i] = check
			}

			_, err := collection.InsertMany(ctx, documents)
			if err != nil {
				fmt.Println(err)
			}
		}

		time.Sleep(1 * time.Minute)
	}
}
