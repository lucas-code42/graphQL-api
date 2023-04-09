package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/lucas-code42/graphql-api/graph"
	"github.com/lucas-code42/graphql-api/mongoDatabase"
	"log"
	"net/http"
)

const defaultPort = "8080"

func main() {
	fmt.Println("antes")
	mongoDatabase.Connect()
	fmt.Println("depois")

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
