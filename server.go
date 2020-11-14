package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/2785/gqlgen-demo/graph"
	"github.com/2785/gqlgen-demo/graph/generated"
	"github.com/2785/gqlgen-demo/graph/model"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/google/uuid"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	testUsers := []*model.User{
		{ID: uuid.New().String(), Age: 20, Name: "Bob"},
		{ID: uuid.New().String(), Age: 18, Name: "Bob's brother Bobby"},
	}

	c := generated.Config{Resolvers: &graph.Resolver{UsersStore: testUsers}}
	c.Directives.Entitlement = func(ctx context.Context, obj interface{}, next graphql.Resolver, resource *string, scope []*string) (res interface{}, err error) {
		return nil, errors.New("access denied")
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
