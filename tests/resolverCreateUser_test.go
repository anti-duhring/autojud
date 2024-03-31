package tests

import (
	"fmt"

	"github.com/99designs/gqlgen/client"
	genGraphql "github.com/anti-duhring/autojud/internal/generated/graphql"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("resolverCustomerRequestAccess", func() {
	type createUserResp struct {
		CreateUser genGraphql.User
	}

	makeRequest := func(options ...client.Option) (createUserResp, error) {
		var resp createUserResp
		err := c.Post(`mutation CreateUser($input: UserInput!) {
			CreateUser(input: $input) {
				id
        name
        email
			} 
		}`,
			&resp,
			options...,
		)
		return resp, err
	}

	It("creates user", func() {
		resp, err := makeRequest()

		fmt.Println(err)
		fmt.Println(resp)
	})
})
