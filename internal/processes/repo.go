package processes

import (
	"context"
	"database/sql"
)

type Repository interface {
	CreateProcessFollow(ctx context.Context, processID string, userID string) (*ProcessFollow, error)
	GetByProcessNumber(ctx context.Context, processNumber string) (*Process, error)
	GetAllByUserID(ctx context.Context, userID string, limit, offset int) ([]*Process, error)
	CountByUserID(ctx context.Context, userID string) (int, error)
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

func (r *RepositoryPostgres) GetAllByUserID(ctx context.Context, userID string, limit, offset int) ([]*Process, error) {
	var processes []*Process

	query := `SELECT p.id, p.process_number, p.court, p.origin, p.judge, p.active_part, p.passive_part, p.created_at, p.updated_at, p.deleted_at FROM processes p JOIN process_follows pf ON p.id = pf.process_id 
  WHERE pf.user_id = $1 
  ORDER BY pf.created_at ASC
  LIMIT $2 
  OFFSET $3;`
	rows, err := r.DB.QueryContext(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var process Process
		err := rows.Scan(&process.ID, &process.ProcessNumber, &process.Court, &process.Origin, &process.Judge, &process.ActivePart, &process.PassivePart, &process.CreatedAt, &process.UpdatedAt, &process.DeletedAt)
		if err != nil {
			return nil, err
		}

		processes = append(processes, &process)
	}

	return processes, nil
}

func (r *RepositoryPostgres) CountByUserID(ctx context.Context, userID string) (int, error) {
	var count int

	query := `SELECT COUNT(*) FROM process_follows WHERE user_id = $1;`
	err := r.DB.QueryRowContext(ctx, query, userID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
