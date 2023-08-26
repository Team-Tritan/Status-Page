package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"tritan.dev/status-page/config"
)

func Init() (*mongo.Client, error) {
	uri := config.LoadConfig().MongoURI

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, nil
}
