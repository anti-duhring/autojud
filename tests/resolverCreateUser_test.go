package tests

import (
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

var _ = Describe("resolverCustomerRequestAccess", func() {
	type createUserResp struct {
		CreateUser genGraphql.User
	}

	makeRequest := func(options ...client.Option) (createUserResp, error) {
		var resp createUserResp
		err := c.Post(`mutation CreateUser($input: CreateUserInput!) {
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

	It("creates a user", func() {
		id := uuid.New()
		input := genGraphql.CreateUserInput{
			Name:     "Matt",
			Email:    "matt@mail.com",
			Password: "password",
		}

		userRepo.(*mocks.MockRepository).Mock.On("Create", mock.Anything, mock.Anything).Return(&user.User{
			ID:        id,
			Name:      input.Name,
			Email:     input.Email,
			Password:  &input.Password,
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
			DeletedAt: nil,
		}, nil)

		resp, err := makeRequest(
			client.Var("input", input),
		)
		Expect(err).ToNot(HaveOccurred())

		Expect(resp.CreateUser.ID).To(Equal(id.String()))
		Expect(resp.CreateUser.Name).To(Equal(input.Name))
		Expect(resp.CreateUser.Email).To(Equal(input.Email))
		Expect(resp.CreateUser.Password).To(BeNil())
		Expect(resp.CreateUser.CreatedAt).ToNot(BeNil())
		Expect(resp.CreateUser.UpdatedAt).ToNot(BeNil())
		Expect(resp.CreateUser.DeletedAt).To(BeNil())
	})

	It("returns an error if email is invalid", func() {
		input := genGraphql.CreateUserInput{
			Name:     "Matt",
			Email:    "mattmail.com",
			Password: "password",
		}

		_, err := makeRequest(
			client.Var("input", input),
		)
		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(ContainSubstring(user.ErrEmailInvalid.Error())))
	})

	It("returns an error if password is empty", func() {
		input := genGraphql.CreateUserInput{
			Name:     "Matt",
			Email:    "matt@mail.com",
			Password: "",
		}

		_, err := makeRequest(
			client.Var("input", input),
		)
		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(ContainSubstring(user.ErrPasswordEmpty.Error())))
	})

	It("returns an error if name is empty", func() {
		input := genGraphql.CreateUserInput{
			Name:     "",
			Email:    "matt@mail.com",
			Password: "password",
		}

		_, err := makeRequest(
			client.Var("input", input),
		)
		Expect(err).To(HaveOccurred())
		Expect(err).To(MatchError(ContainSubstring(user.ErrNameEmpty.Error())))
	})

})
