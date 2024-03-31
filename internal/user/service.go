package user

import "context"

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

	return s.Repository.Create(ctx, u)
}
