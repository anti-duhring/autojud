package cron

import (
	"context"
	"database/sql"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*Cron, error)
	LockCron(ctx context.Context, crawlerId CRAWLER_ID, tx *sql.Tx) error
	WithTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error
}

type RepositoryPostgres struct {
	DB *sql.DB
}

func NewRepositoryPostgres(DB *sql.DB) *RepositoryPostgres {
	return &RepositoryPostgres{DB: DB}
}

func (r *RepositoryPostgres) GetAll(ctx context.Context) ([]*Cron, error) {
	var crons []*Cron

	query := `SELECT id, crawler_id, created_at, last_run, enabled FROM crons SKIP LOCKED;`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cron Cron
		err := rows.Scan(&cron.ID, &cron.CrawlerID, &cron.CreatedAt, &cron.LastRun, &cron.Enabled)
		if err != nil {
			return nil, err
		}
		crons = append(crons, &cron)
	}

	return crons, nil
}

func (r *RepositoryPostgres) LockCron(ctx context.Context, crawlerId CRAWLER_ID, tx *sql.Tx) error {
	query := `SELECT id, crawler_id, created_at, last_run, enabled FROM crons WHERE crawler_id = $1 FOR UPDATE;`
	_, err := tx.ExecContext(ctx, query, crawlerId)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryPostgres) WithTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := fn(tx); err != nil {
		return err
	}

	return tx.Commit()
}
