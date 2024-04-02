package user

import (
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, user User) (*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

type RepositoryPostgres struct {
	DB *sql.DB
}

func NewRepositoryPostgres(DB *sql.DB) *RepositoryPostgres {
	return &RepositoryPostgres{DB: DB}
}

func (r *RepositoryPostgres) Create(ctx context.Context, user User) (*User, error) {
	var createdUser User

	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING  id, name, email, password, created_at, updated_at, deleted_at;`
	err := r.DB.QueryRowContext(ctx, query, user.Name, user.Email, user.Password).Scan(&createdUser.ID, &createdUser.Name, &createdUser.Email, &createdUser.Password, &createdUser.CreatedAt, &createdUser.UpdatedAt, &createdUser.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}

func (r *RepositoryPostgres) GetByID(ctx context.Context, id string) (*User, error) {
	var user User

	query := `SELECT * FROM users WHERE id = $1;`
	err := r.DB.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *RepositoryPostgres) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user User

	query := `SELECT * FROM users WHERE email = $1;`
	err := r.DB.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
