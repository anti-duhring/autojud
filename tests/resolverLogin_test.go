package tests

import (
	"database/sql"
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

var _ = Describe("resolverLogin", func() {
	type loginResp struct {
		Login genGraphql.AuthResponse
	}

	makeRequest := func(options ...client.Option) (loginResp, error) {
		var resp loginResp
		err := c.Post(`mutation Login($email: String!, $password: String!) {
			Login(email: $email, password: $password) {
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

	It("authenticates a user", func() {
		id := uuid.New()
		name := "Matt"
		email := "matt@mail.com"
		password := "password"
		encryptedPassword, _ := users.HashPassword(password)

		userRepo.(*mocks.MockRepository).Mock.On("GetByEmail", mock.Anything, mock.Anything).Return(&users.User{
			ID:        id,
			Name:      name,
			Email:     email,
			Password:  &encryptedPassword,
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
			DeletedAt: nil,
		}, nil)

		resp, err := makeRequest(
			client.Var("email", email),
			client.Var("password", password),
		)
		Expect(err).ToNot(HaveOccurred())

		Expect(resp.Login.User.ID).To(Equal(id.String()))
		Expect(resp.Login.User.Name).To(Equal(name))
		Expect(resp.Login.User.Email).To(Equal(email))
		Expect(resp.Login.User.Password).To(BeNil())
		Expect(resp.Login.User.CreatedAt).ToNot(BeNil())
		Expect(resp.Login.User.UpdatedAt).ToNot(BeNil())
		Expect(resp.Login.User.DeletedAt).To(BeNil())

		Expect(resp.Login.Token).ToNot(BeNil())
		Expect(resp.Login.TokenExp).ToNot(BeNil())
	})

	It("throws an error if password is wrong", func() {
		id := uuid.New()
		name := "Matt"
		email := "matt@mail.com"
		password := "wrongpassword"
		encryptedPassword, _ := users.HashPassword("password")

		userRepo.(*mocks.MockRepository).Mock.On("GetByEmail", mock.Anything, mock.Anything).Return(&users.User{
			ID:        id,
			Name:      name,
			Email:     email,
			Password:  &encryptedPassword,
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
			DeletedAt: nil,
		}, nil)

		_, err := makeRequest(
			client.Var("email", email),
			client.Var("password", password),
		)
		Expect(err).To(HaveOccurred())

		Expect(err).To(MatchError(ContainSubstring(users.ErrInvalidCredentials.Error())))
	})

	It("returns an error if the user does not exist", func() {
		email := "matt@mail.com"
		password := "password"

		userRepo.(*mocks.MockRepository).Mock.On("GetByEmail", mock.Anything, mock.Anything).Return(nil, sql.ErrNoRows)

		_, err := makeRequest(
			client.Var("email", email),
			client.Var("password", password),
		)
		Expect(err).To(HaveOccurred())

		Expect(err).To(MatchError(ContainSubstring(users.ErrInvalidCredentials.Error())))
	})
})
