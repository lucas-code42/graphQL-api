package mongoDatabase

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Replace the placeholder with your Atlas connection string
const uri = "mongodb://root:example@localhost:27017/?connect=direct"

func Connect() (*mongo.Client, error) {
	dbURL := os.Getenv("MONGO_URL")
	fmt.Println(dbURL)

	// Set client options
	clientOptions := options.Client()
	clientOptions.SetDirect(true).ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("** 1 **", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal("** 2 **", err)
	}
	// Return successful connection
	return client, nil

}
