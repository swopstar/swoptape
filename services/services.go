// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package services

import (
	"context"
	"errors"
	"log/slog"

	"github.com/swopstar/gokit/jobs"
	"github.com/swopstar/swoptape/config"
	"github.com/swopstar/swoptape/database"
	"github.com/swopstar/swoptape/services/identity"
)

type Services struct {
	Jobs     *jobs.JobService
	Identity identity.IService
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

	identitySvc, err := identity.New(ctx, cfg, l.WithGroup("identity"), db)
	if err != nil {
		return nil, svcInitError(l, "identity", err)
	}

	return &Services{
		Jobs:     jobSvc,
		Identity: identitySvc,
	}, nil
}

func (svc *Services) Close() error {
	return svc.Jobs.Shutdown()
}

func svcInitError(l *slog.Logger, name string, err error) error {
	l.Error("Failed to initialise service", "service", name, "err", err)
	return errors.Join(&ServiceInitError{
		ServiceName: name,
	}, err)
}
