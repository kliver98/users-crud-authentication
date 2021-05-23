package database

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetConnection() (*mongo.Client, error) {
	uri := os.Getenv("MONGODB_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client, err

}

func IsDown(client *mongo.Client, ctx context.Context) bool {
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return true
	}
	return false
}
