package cron

import (
	"context"
	"database/sql"
	"time"

	crawjud "github.com/anti-duhring/crawjud/pkg/crawler"
)

type CRAWLER_ID string

type CRAWLER_FUNCTION func(*Service, *sql.Tx) error

var crawlers = map[CRAWLER_ID]CRAWLER_FUNCTION{
	"TJPE_MOVIMENTACAO": TJPEMoves,
}

func Crawl(crawlerID CRAWLER_ID, s *Service) error {
	return s.ExecuteCron(crawlerID, TJPEMoves, context.Background())
}

func TJPEMoves(s *Service, tx *sql.Tx) error {
	processes, err := crawjud.TJPE()
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	for number, development := range processes {
		_, _, err := s.ProcessService.CreateDevelopment(number, now.String(), development, context.Background())
		if err != nil {
			return err
		}
	}

	return nil
}
