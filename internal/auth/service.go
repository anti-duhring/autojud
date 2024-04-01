package auth

import (
	"context"

	"github.com/anti-duhring/autojud/internal/user"
	"github.com/anti-duhring/autojud/pkg/jwt"
)

type Service struct {
	UserService user.Service
}

func NewService(s user.Service) *Service {
	return &Service{UserService: s}
}

func (s *Service) Register(ctx context.Context, user user.User) (*Response, error) {
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
