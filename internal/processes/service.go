package processes

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

func (s *Service) FollowProcess(processID uuid.UUID, userID uuid.UUID, ctx context.Context) (*ProcessFollow, error) {
	processFollow, err := s.Repository.CreateProcessFollow(ctx, processID.String(), userID.String())
	if err != nil {
		logger.Error("error following process", err)
		return nil, ErrInternal
	}

	return processFollow, nil
}

func (s *Service) GetByProcessNumber(processNumber string, ctx context.Context) (*Process, error) {
	process, err := s.Repository.GetByProcessNumber(ctx, processNumber)
	if err != nil {
		logger.Error("error getting process by process number", err)
		return nil, ErrInternal
	}

	return process, nil
}

func (s *Service) GetProcessFromUser(userID uuid.UUID, limit, offset int, ctx context.Context) ([]*Process, error) {
	processes, err := s.Repository.GetAllByUserID(ctx, userID.String(), limit, offset)
	if err != nil {
		logger.Error("error getting process from user", err)
		return nil, ErrInternal
	}

	return processes, nil
}

func (s *Service) CountProcessFromUser(userID uuid.UUID, ctx context.Context) (int, error) {
	count, err := s.Repository.CountByUserID(ctx, userID.String())
	if err != nil {
		logger.Error("error counting process from user", err)
		return 0, ErrInternal
	}

	return count, nil
}
