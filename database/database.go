package database

import (
	"log/slog"

	gkdatabase "github.com/swopstar/gokit/database"
)

type Database struct {
	*gkdatabase.Database
}

func Open(cfg *gkdatabase.Config, l *slog.Logger) (*Database, error) {
	db, err := gkdatabase.Open(cfg, l)
	if err != nil {
		return nil, err
	}

	db.AddModels(Models...)

	return &Database{Database: db}, nil
}
