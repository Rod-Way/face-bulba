package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	CO "faceBulba/config"
)

func GetDB(collection string) (*mongo.Client, *mongo.Collection, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to  MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(CO.GetDBPath()))
	if err != nil {
		cancel()
		return nil, nil, nil, nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	col := client.Database(os.Getenv("DB_NAME")).Collection(collection)
	return client, col, ctx, cancel, nil
}
