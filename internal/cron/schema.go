package cron

import "github.com/google/uuid"

type Cron struct {
	ID        uuid.UUID  `json:"id"`
	CrawlerID CRAWLER_ID `json:"crawler_id"`
	CreatedAt string     `json:"created_at"`
	LastRun   string     `json:"last_run"`
	Enabled   bool       `json:"enabled"`
}
