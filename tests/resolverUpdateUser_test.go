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

var _ = Describe("resolverUpdateUser", func() {
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
		oldEncryptedPassword, _ := user.HashPassword("oldpassword")

		input := &genGraphql.UpdateUserInput{
			Name:     &name,
			Email:    &email,
			Password: &password,
		}

		userRepo.(*mocks.MockRepository).Mock.On("GetByID", mock.Anything, mock.Anything).Return(&user.User{
			ID:        id,
			Name:      "Tom",
			Email:     "tombrady@nepatrios.com",
			Password:  &oldEncryptedPassword,
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
			DeletedAt: nil,
		}, nil)

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
			mocks.MockToken(),
			mocks.MockUserID(id.String()),
			client.Var("input", input),
		)
		Expect(err).ToNot(HaveOccurred())

		Expect(resp.UpdateUser.ID).To(Equal(id.String()))
		Expect(resp.UpdateUser.Name).To(Equal(name))
		Expect(resp.UpdateUser.Email).To(Equal(email))
	})
	FIt("updates only the user's email", func() {
		id := uuid.New()
		name := "Matt"
		newEmail := "matt@mail.com"
		encryptedPassword, _ := user.HashPassword("password")

		input := &genGraphql.UpdateUserInput{
			Email: &newEmail,
		}

		userRepo.(*mocks.MockRepository).Mock.On("GetByID", mock.Anything, mock.Anything).Return(&user.User{
			ID:        id,
			Name:      name,
			Email:     "tombrady@nepatrios.com",
			Password:  &encryptedPassword,
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
			DeletedAt: nil,
		}, nil)

		userRepo.(*mocks.MockRepository).Mock.On("Update", mock.Anything, mock.Anything).Return(&user.User{
			ID:        id,
			Name:      name,
			Email:     newEmail,
			Password:  &encryptedPassword,
			CreatedAt: time.Now().String(),
			UpdatedAt: time.Now().String(),
			DeletedAt: nil,
		}, nil)

		resp, err := makeRequest(
			mocks.MockToken(),
			mocks.MockUserID(id.String()),
			client.Var("input", input),
		)
		Expect(err).ToNot(HaveOccurred())

		Expect(resp.UpdateUser.ID).To(Equal(id.String()))
		Expect(resp.UpdateUser.Name).To(Equal(name))
		Expect(resp.UpdateUser.Email).To(Equal(newEmail))
	})

})
