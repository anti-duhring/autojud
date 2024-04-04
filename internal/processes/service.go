package processes

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	Repository Repository
}

func NewService(r Repository) *Service {
	return &Service{Repository: r}
}

func (s *Service) FollowProcess(processID uuid.UUID, userID uuid.UUID, ctx context.Context) (*ProcessFollow, error) {
	return s.Repository.CreateProcessFollow(ctx, processID.String(), userID.String())
}
