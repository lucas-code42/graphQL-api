package mongoDatabase

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TestAPI struct {
	_Id                 string `bson:"_id"`
	Name                string `bson:"name"`
	ProgrammingLanguage string `bson:"programmingLanguage"`
}

const uri = "mongodb://root:example@localhost:27017/?connect=direct"

func connect() (*mongo.Client, error) {
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

func GetAll() ([]TestAPI, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := connect()
	if err != nil {
		log.Fatal("** 3 **", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("account").Collection("test_api")
	filter := bson.M{}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal("** 4 **", err)
	}
	defer cur.Close(ctx)

	var testAPI []TestAPI
	for cur.Next(ctx) {
		var tmp TestAPI
		if err := cur.Decode(&tmp); err != nil {
			log.Fatal("** 5 **", err)
		}

		testAPI = append(testAPI, tmp)
	}

	if err := cur.Err(); err != nil {
		log.Fatal("** 6 **", err)
	}

	//for _, v := range testAPI {
	//	fmt.Println(v.Name)
	//	fmt.Println(v.ProgrammingLanguage)
	//}

	return testAPI, nil

}
