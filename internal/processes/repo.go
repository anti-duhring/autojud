package processes

import (
	"context"
	"database/sql"
)

type Repository interface {
	CreateProcessFollow(ctx context.Context, processID string, userID string) (*ProcessFollow, error)
	GetByProcessNumber(ctx context.Context, processNumber string) (*Process, error)
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

func (r *RepositoryPostgres) GetByProcessNumber(ctx context.Context, processNumber string) (*Process, error) {
	var process Process

	query := `SELECT id, process_number, court, origin, judge, active_part, passive_part, created_at, updated_at, deleted_at FROM processes WHERE process_number = $1;`
	err := r.DB.QueryRowContext(ctx, query, processNumber).Scan(&process.ID, &process.ProcessNumber, &process.Court, &process.Origin, &process.Judge, &process.ActivePart, &process.PassivePart, &process.CreatedAt, &process.UpdatedAt, &process.DeletedAt)
	if err != nil {
		return nil, err
	}

	return &process, nil
}
