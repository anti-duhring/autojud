package tests

import (
	"time"

	"github.com/99designs/gqlgen/client"
	genGraphql "github.com/anti-duhring/autojud/internal/generated/graphql"
	"github.com/anti-duhring/autojud/internal/users"
	"github.com/anti-duhring/autojud/tests/mocks"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("resolverRegister", func() {
	type registerResp struct {
		Register genGraphql.AuthResponse
	}

	makeRequest := func(options ...client.Option) (registerResp, error) {
		var resp registerResp
		err := c.Post(`mutation Register($input: CreateUserInput!) {
			Register(input: $input) {
				user {
          id
          name
          email
        }
        token
        tokenExp
			} 
		}`,
			&resp,
			options...,
		)
		return resp, err
	}

	It("register a user", func() {
		id := uuid.New()
		input := genGraphql.CreateUserInput{
			Name:     "Matt",
			Email:    "matt@mail.com",
			Password: "password",
		}

		userRepo.(*mocks.MockRepository).Mock.On("Create", mock.Anything, mock.Anything).Return(&users.User{
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

		Expect(resp.Register.User.ID).To(Equal(id.String()))
		Expect(resp.Register.User.Name).To(Equal(input.Name))
		Expect(resp.Register.User.Email).To(Equal(input.Email))
		Expect(resp.Register.User.Password).To(BeNil())
		Expect(resp.Register.User.CreatedAt).ToNot(BeNil())
		Expect(resp.Register.User.UpdatedAt).ToNot(BeNil())
		Expect(resp.Register.User.DeletedAt).To(BeNil())

		Expect(resp.Register.Token).ToNot(BeNil())
		Expect(resp.Register.TokenExp).ToNot(BeNil())
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
		Expect(err).To(MatchError(ContainSubstring(users.ErrEmailInvalid.Error())))
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
		Expect(err).To(MatchError(ContainSubstring(users.ErrPasswordEmpty.Error())))
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
		Expect(err).To(MatchError(ContainSubstring(users.ErrNameEmpty.Error())))
	})

})
