package database

import (
	"log/slog"

	gkdatabase "github.com/swopstar/gokit/database"
	"github.com/swopstar/gokit/jobs"
)

type Database struct {
	*gkdatabase.Database
}

func Open(cfg *gkdatabase.Config, l *slog.Logger) (*Database, error) {
	db, err := gkdatabase.Open(cfg, l)
	if err != nil {
		return nil, err
	}

	db.AddModels(jobs.Models...)
	db.AddModels(Models...)

	return &Database{Database: db}, nil
}
