package auth

import (
	"context"

	"github.com/anti-duhring/autojud/internal/users"
	"github.com/anti-duhring/autojud/pkg/jwt"
)

type Service struct {
	UserService users.Service
}

func NewService(s users.Service) *Service {
	return &Service{UserService: s}
}

func (s *Service) Register(ctx context.Context, user users.User) (*Response, error) {
	createdUser, err := s.UserService.Create(user, ctx)
	if err != nil {
		return nil, err
	}

	token, exp, err := jwt.GenerateToken(createdUser.ID.String())
	if err != nil {
		return nil, err
	}

	return &Response{
		User:     *createdUser,
		Token:    token,
		TokenExp: exp,
	}, nil
}

func (s *Service) Login(ctx context.Context, email, password string) (*Response, error) {
	loggedUser, err := s.UserService.GetByEmail(email, ctx)
	if err != nil {
		return nil, err
	}

	if !users.CheckPasswordHash(password, *loggedUser.Password) {
		return nil, users.ErrInvalidCredentials
	}

	token, exp, err := jwt.GenerateToken(loggedUser.ID.String())
	if err != nil {
		return nil, err
	}

	return &Response{
		User:     *loggedUser,
		Token:    token,
		TokenExp: exp,
	}, nil
}
