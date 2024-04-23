package cron

import (
	"context"
)

type Service struct {
	Repository Repository
}

func NewService(r Repository) *Service {
	return &Service{Repository: r}
}

func (s *Service) GetAllCrons(ctx context.Context) ([]*Cron, error) {
	crons, err := s.Repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return crons, nil
}
