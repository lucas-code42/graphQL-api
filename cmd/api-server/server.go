package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/lucas-code42/graphql-api/graph"
	"github.com/lucas-code42/graphql-api/mongoDatabase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"
const uri = "mongodb://root:example@localhost:27017/?connect=direct"

func connect() (*mongo.Client, error) {
	dbURL := os.Getenv("MONGO_URL")
	fmt.Println(dbURL)

	clientOptions := options.Client()
	clientOptions.SetDirect(true).ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("** 1 **", err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal("** 2 **", err)
	}

	return client, nil
}

func main() {
	c, err := connect()
	if err != nil {
		log.Fatal("Erro ao se conectar com MongoDB", err)
	}
	defer c.Disconnect(context.Background())

	ac := mongoDatabase.InitMongo(c)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Account: ac,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
