package user

import (
	"context"
)

type RepositoryMocked struct{}

func NewRepositoryMocked() *RepositoryMocked {
	return &RepositoryMocked{}
}

func (r *RepositoryMocked) Create(ctx context.Context, user User) (User, error) {
	return User{}, nil
}
