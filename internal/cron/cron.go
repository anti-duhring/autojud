package cron

import (
	"context"
	"fmt"
	"time"

	"github.com/anti-duhring/goncurrency/pkg/logger"
	"github.com/go-co-op/gocron"
)

func Init(s *Service) {
	scheduler := gocron.NewScheduler(time.UTC)

	_, err := scheduler.Every(1).Minute().Do(func() {
		if err := run(s); err != nil {
			logger.Error("error running cron", err)
		}
	})
	if err != nil {
		logger.Error("error scheduling cron", err)
	}

	scheduler.StartAsync()
}

func run(s *Service) error {
	crons, err := s.GetAllCrons(context.Background())
	if err != nil {
		logger.Error("error getting all crons", err)
		return err
	}

	for _, cron := range crons {
		logger.Debug(fmt.Sprintf("executing cron %s", cron.CrawlerID))

		if err := s.ExecuteCron(cron.CrawlerID, crawlers[cron.CrawlerID], context.Background()); err != nil {
			logger.Error(fmt.Sprintf("error executing cron %s", cron.CrawlerID), err)
		}
	}

	return nil
}
