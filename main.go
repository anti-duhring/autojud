package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/anti-duhring/autojud/internal/auth"
	"github.com/anti-duhring/autojud/internal/db"
	genGraphql "github.com/anti-duhring/autojud/internal/generated/graphql"
	"github.com/anti-duhring/autojud/internal/graphql/resolvers"
	"github.com/anti-duhring/autojud/internal/user"
	"github.com/anti-duhring/goncurrency/pkg/logger"
	"github.com/go-chi/chi"
)

func main() {
	db, err := db.Init()
	if err != nil {
		panic(err)
	}

	userRepo := user.NewRepositoryPostgres(db)
	userService := user.NewService(userRepo)
	authService := auth.NewService(*userService)

	router := chi.NewRouter()
	router.Use(auth.Middleware(*userService))

	srv := handler.NewDefaultServer(genGraphql.NewExecutableSchema(genGraphql.Config{Resolvers: &resolvers.Resolver{
		UserService: userService,
		AuthService: authService,
	}}))

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	logger.Debug("Running server on port " + port)
	router.Handle("/query", srv)

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
