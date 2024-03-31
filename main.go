package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/anti-duhring/autojud/internal/db"
	genGraphql "github.com/anti-duhring/autojud/internal/generated/graphql"
	"github.com/anti-duhring/autojud/internal/graphql/resolvers"
	"github.com/anti-duhring/autojud/internal/user"
	"github.com/anti-duhring/goncurrency/pkg/logger"
)

var userService *user.Service

func main() {
	db, err := db.Init()
	if err != nil {
		panic(err)
	}

	userRepo := user.NewRepositoryPostgres(db)
	userService = user.NewService(userRepo)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	srv := handler.NewDefaultServer(genGraphql.NewExecutableSchema(genGraphql.Config{Resolvers: &resolvers.Resolver{}}))

	logger.Debug(fmt.Sprintf("connect to http://localhost:%s/query for GraphQL", port))
	http.Handle("/query", srv)
	http.ListenAndServe(":"+port, nil)
}
