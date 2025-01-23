package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoUrl = "mongodb://mongo:27017"
	gRpcPort = "50001"
)

var client *mongo.Client

type Config struct {
}

func main() {
	mongoClient, err := connectToMongo(mongoUrl)
	if err != nil {
		log.Panic(err)
	}
	client = mongoClient
}

func connectToMongo(url string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(url)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to mongo: %w", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, fmt.Errorf("unable to ping mongo: %w", err)
	}

	return client, nil
}
