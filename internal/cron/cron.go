package cron

import (
	"context"

	"github.com/anti-duhring/goncurrency/pkg/logger"
)

func Run(s *Service) error {
	crons, err := s.GetAllCrons(context.Background())
	if err != nil {
		logger.Error("error getting all crons", err)
		return err
	}

	for _, cron := range crons {

	}

	return nil
}
