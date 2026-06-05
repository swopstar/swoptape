package services

import (
	"context"
	"log/slog"

	"github.com/swopstar/swoptape/config"
	"github.com/swopstar/swoptape/database"
)

type Services struct {
	logger *slog.Logger
}

func New(
	ctx context.Context,
	cfg *config.Config,
	l *slog.Logger,
	db *database.Database,
) (*Services, error) {
	panic("TODO")
}

func (svc *Services) Close() error {
	panic("TODO")
}
