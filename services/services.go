package services

import (
	"context"
	"errors"
	"log/slog"

	"github.com/swopstar/gokit/jobs"
	"github.com/swopstar/swoptape/config"
	"github.com/swopstar/swoptape/database"
)

type Services struct {
	Jobs *jobs.JobService
}

func New(
	ctx context.Context,
	cfg *config.Config,
	l *slog.Logger,
	db *database.Database,
) (*Services, error) {
	jobSvc, err := jobs.NewJobService(ctx, &cfg.Jobs, l.WithGroup("jobs"), db.Gorm)
	if err != nil {
		return nil, svcInitError(l, "jobs", err)
	}

	return &Services{
		Jobs: jobSvc,
	}, nil
}

func (svc *Services) Close() error {
	panic("TODO")
}

func svcInitError(l *slog.Logger, name string, err error) error {
	l.Error("Failed to initialise service", "service", name, "err", err)
	return errors.Join(&ServiceInitError{
		ServiceName: name,
	}, err)
}
