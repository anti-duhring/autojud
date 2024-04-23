package cron

import (
	"context"
	"database/sql"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*Cron, error)
}

type RepositoryPostgres struct {
	DB *sql.DB
}

func NewRepositoryPostgres(DB *sql.DB) *RepositoryPostgres {
	return &RepositoryPostgres{DB: DB}
}

func (r *RepositoryPostgres) GetAll(ctx context.Context) ([]*Cron, error) {
	var crons []*Cron

	query := `SELECT id, crawler_id, created_at, last_run, enabled FROM crons;`
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
