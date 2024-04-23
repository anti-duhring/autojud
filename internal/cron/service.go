package cron

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/anti-duhring/autojud/internal/processes"
	"github.com/anti-duhring/goncurrency/pkg/logger"
)

type Service struct {
	Repository     Repository
	ProcessService processes.Service
}

func NewService(r Repository, p processes.Service) *Service {
	return &Service{Repository: r, ProcessService: p}
}

func (s *Service) GetAllCrons(ctx context.Context) ([]*Cron, error) {
	crons, err := s.Repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return crons, nil
}

func (s *Service) ExecuteCron(cID CRAWLER_ID, c CRAWLER_FUNCTION, ctx context.Context) error {
	return s.Repository.WithTransaction(ctx, func(tx *sql.Tx) error {
		err := s.Repository.LockCron(ctx, cID, tx)
		if err != nil {
			logger.Error("error locking cron", err)
			return err
		}

		err = c(s, tx)
		if err != nil {
			logger.Error(fmt.Sprintf("error executing cron %s", cID), err)
			return err
		}

		return nil
	})
}
