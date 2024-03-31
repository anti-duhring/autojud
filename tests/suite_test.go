package tests

import (
	"os"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	genGraphql "github.com/anti-duhring/autojud/internal/generated/graphql"
	"github.com/anti-duhring/autojud/internal/graphql/resolvers"
	"github.com/anti-duhring/autojud/internal/user"
	"github.com/anti-duhring/autojud/tests/mocks"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var c *client.Client

func Test(t *testing.T) {
	os.Setenv("TZ", "UTC")

	RegisterFailHandler(Fail)
	RunSpecs(t, "Test suite")
}

var _ = BeforeSuite(func() {
	t := GinkgoT()
	userRepo := mocks.NewMockRepository(t)
	userService := user.NewService(userRepo)

	server := handler.NewDefaultServer(genGraphql.NewExecutableSchema(genGraphql.Config{Resolvers: &resolvers.Resolver{
		UserService: userService,
	}}))
	c = client.New(server)
})