package processes

import (
	"context"
	"database/sql"
	"errors"

	crawjud "github.com/anti-duhring/crawjud/pkg/utils"
	"github.com/anti-duhring/goncurrency/pkg/logger"
	"github.com/google/uuid"
)

type Service struct {
	Repository Repository
}

func NewService(r Repository) *Service {
	return &Service{Repository: r}
}

func (s *Service) FollowProcess(processID uuid.UUID, userID uuid.UUID, ctx context.Context) (*ProcessFollow, error) {
	processFollow, err := s.Repository.CreateProcessFollow(ctx, processID.String(), userID.String())
	if err != nil {
		logger.Error("error following process", err)
		return nil, ErrInternal
	}

	return processFollow, nil
}

func (s *Service) GetByProcessNumber(processNumber string, ctx context.Context) (*Process, error) {
	process, err := s.Repository.GetByProcessNumber(ctx, processNumber)
	if err != nil {
		logger.Error("error getting process by process number", err)
		return nil, err
	}

	return process, nil
}

func (s *Service) GetProcessFromUser(userID uuid.UUID, limit, offset int, ctx context.Context) ([]*Process, error) {
	processes, err := s.Repository.GetAllByUserID(ctx, userID.String(), limit, offset)
	if err != nil {
		logger.Error("error getting process from user", err)
		return nil, ErrInternal
	}

	return processes, nil
}

func (s *Service) CountProcessFromUser(userID uuid.UUID, ctx context.Context) (int, error) {
	count, err := s.Repository.CountByUserID(ctx, userID.String())
	if err != nil {
		logger.Error("error counting process from user", err)
		return 0, ErrInternal
	}

	return count, nil
}

func (s *Service) CreatePendingProcess(processNumber string, ctx context.Context) (*PendingProcess, *Process, error) {
	var court Court

	courtCode, err := crawjud.GetCourtByProcessNumber(processNumber)
	if err != nil || courtCode == nil {
		logger.Error("error getting court by process number", err)
		court = COURT_UNKNOWN
	}

	if courtCode != nil {
		court = parseCourt(*courtCode)
	}
	process := &Process{
		ProcessNumber: processNumber,
		Court:         court,
	}

	createdProcess, err := s.Repository.CreateProcess(ctx, process)
	if err != nil {
		logger.Error("error creating process", err)
		return nil, nil, ErrInternal
	}

	pendingProcess, err := s.Repository.CreatePendingProcess(ctx, createdProcess.ID.String())
	if err != nil {
		logger.Error("error creating pending process", err)
		return nil, nil, ErrInternal
	}

	return pendingProcess, createdProcess, nil
}

func (s *Service) CreateDevelopment(processNumber string, developmentDate, description string, ctx context.Context) (*ProcessDevelopment, *Process, error) {
	var process *Process
	var err error

	process, err = s.Repository.GetByProcessNumber(ctx, processNumber)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		_, process, err = s.CreatePendingProcess(processNumber, ctx)
		if err != nil {
			return nil, nil, err
		}
	}

	processDevelopment, err := s.Repository.CreateProcessDevelopment(ctx, process.ID.String(), developmentDate, description)
	if err != nil {
		logger.Error("error creating process development", err)
		return nil, nil, ErrInternal
	}

	return processDevelopment, process, nil
}
