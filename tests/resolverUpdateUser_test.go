package tests

import (
	"fmt"
	"time"

	"github.com/99designs/gqlgen/client"
	genGraphql "github.com/anti-duhring/autojud/internal/generated/graphql"
	"github.com/anti-duhring/autojud/internal/user"
	"github.com/anti-duhring/autojud/tests/mocks"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = FDescribe("resolverUpdateUser", func() {
	type resp struct {
		UpdateUser genGraphql.User
	}

	makeRequest := func(options ...client.Option) (resp, error) {
		var resp resp
		err := c.Post(`mutation UpdateUser($input: UpdateUserInput!) {
			UpdateUser(input: $input) {
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

	It("updates a user", func() {
		id := uuid.New()
		name := "Matt"
		email := "matt@mail.com"
		password := "password"
		encryptedPassword, _ := user.HashPassword(password)

		input := &genGraphql.UpdateUserInput{
			Name:     &name,
			Email:    &email,
			Password: &password,
		}

		userRepo.(*mocks.MockRepository).Mock.On("Update", mock.Anything, mock.Anything).Return(&user.User{
			ID:        id,
			Name:      name,
			Email:     email,
			Password:  &encryptedPassword,
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
			DeletedAt: nil,
		}, nil)

		resp, err := makeRequest(
			client.Var("input", input),
		)
		Expect(err).ToNot(HaveOccurred())

		fmt.Println(resp)
	})

})
