package tests

import (
	"os"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/anti-duhring/autojud/internal/auth"
	genGraphql "github.com/anti-duhring/autojud/internal/generated/graphql"
	"github.com/anti-duhring/autojud/internal/graphql/resolvers"
	"github.com/anti-duhring/autojud/internal/user"
	"github.com/anti-duhring/autojud/tests/mocks"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var c *client.Client
var server *handler.Server
var userRepo user.Repository

func Test(t *testing.T) {
	os.Setenv("TZ", "UTC")

	RegisterFailHandler(Fail)
	RunSpecs(t, "Test suite")
}

var _ = BeforeSuite(func() {
	t := GinkgoT()
	userRepo = mocks.NewMockRepository(t)
	userService := user.NewService(userRepo)
	authService := auth.NewService(*userService)

	config := genGraphql.Config{Resolvers: &resolvers.Resolver{
		UserService: userService,
		AuthService: authService,
	}}
	config.Directives.Auth = auth.AuthDirective

	server = handler.NewDefaultServer(genGraphql.NewExecutableSchema(config))
	c = client.New(server)
})
