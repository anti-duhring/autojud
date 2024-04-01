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

	hashedPassword, err := hashPassword(*u.Password)
	if err != nil {
		logger.Error("error hashing password", err)
		return nil, err
	}

	u.Password = &hashedPassword

	return s.Repository.Create(ctx, u)
}

func (s *Service) GetByID(id uuid.UUID, ctx context.Context) (*User, error) {
	return s.Repository.GetByID(ctx, id.String())
}
