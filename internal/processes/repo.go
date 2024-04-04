package processes

import (
	"context"
	"database/sql"
)

type Repository interface {
	CreateProcessFollow(ctx context.Context, processID string, userID string) (*ProcessFollow, error)
}

type RepositoryPostgres struct {
	DB *sql.DB
}

func NewRepositoryPostgres(DB *sql.DB) *RepositoryPostgres {
	return &RepositoryPostgres{DB: DB}
}

func (r *RepositoryPostgres) CreateProcessFollow(ctx context.Context, processID string, userID string) (*ProcessFollow, error) {
	var createdProcessFollow ProcessFollow

	query := `INSERT INTO process_follows (process_id, user_id) VALUES ($1, $2) RETURNING id, process_id, user_id, created_at,  deleted_at;`
	err := r.DB.QueryRowContext(ctx, query, processID, userID).Scan(&createdProcessFollow.ID, &createdProcessFollow.ProcessID, &createdProcessFollow.UserID, &createdProcessFollow.CreatedAt, &createdProcessFollow.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &createdProcessFollow, nil

}
