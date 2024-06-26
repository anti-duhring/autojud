package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/anti-duhring/autojud/internal/auth"
	"github.com/anti-duhring/autojud/internal/cron"
	"github.com/anti-duhring/autojud/internal/db"
	genGraphql "github.com/anti-duhring/autojud/internal/generated/graphql"
	"github.com/anti-duhring/autojud/internal/graphql/resolvers"
	"github.com/anti-duhring/autojud/internal/processes"
	"github.com/anti-duhring/autojud/internal/users"
	"github.com/anti-duhring/goncurrency/pkg/logger"
	"github.com/go-chi/chi"
)

func main() {
	db, err := db.Init()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := users.NewRepositoryPostgres(db)
	processRepo := processes.NewRepositoryPostgres(db)
	cronRepo := cron.NewRepositoryPostgres(db)

	userService := users.NewService(userRepo)
	processService := processes.NewService(processRepo)
	authService := auth.NewService(*userService)
	cronService := cron.NewService(cronRepo, *processService)

	cron.Init(cronService)

	router := chi.NewRouter()
	router.Use(auth.Middleware(*userService))

	c := genGraphql.Config{Resolvers: &resolvers.Resolver{
		UserService:    userService,
		AuthService:    authService,
		ProcessService: processService,
	}}
	c.Directives.Auth = auth.AuthDirective

	srv := handler.NewDefaultServer(genGraphql.NewExecutableSchema(c))

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
