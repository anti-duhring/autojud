package user

import (
	"context"

	"github.com/anti-duhring/goncurrency/pkg/logger"
	"github.com/google/uuid"
)

type Service struct {
	Repository Repository
}

func NewService(r Repository) *Service {
	return &Service{Repository: r}
}

func (s *Service) Create(u User, ctx context.Context) (*User, error) {
	err := u.Validate()
	if err != nil {
		return nil, err
	}

	hashedPassword, err := HashPassword(*u.Password)
	if err != nil {
		logger.Error("error hashing password", err)
		return nil, err
	}

	u.Password = &hashedPassword

	createdUser, err := s.Repository.Create(ctx, u)
	if err != nil {
		logger.Error("error creating user", err)
		return nil, ErrInternal
	}

	return createdUser, nil
}

func (s *Service) GetByID(id uuid.UUID, ctx context.Context) (*User, error) {
	return s.Repository.GetByID(ctx, id.String())
}

func (s *Service) GetByEmail(email string, ctx context.Context) (*User, error) {
	user, err := s.Repository.GetByEmail(ctx, email)
	if err != nil {
		logger.Error("error getting user by email", err)
		return nil, ErrInvalidCredentials
	}

	return user, nil
}
