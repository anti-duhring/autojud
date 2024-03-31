package user

import "context"

type Service struct {
	Repository Repository
}

func NewService(r Repository) *Service {
	return &Service{Repository: r}
}

func (s *Service) Create(u User, ctx context.Context) (*User, error) {
	return s.Repository.Create(ctx, u)
}
