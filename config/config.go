// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package config

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/pelletier/go-toml/v2"
	"github.com/swopstar/gokit/config"
	"github.com/swopstar/gokit/database"
	"github.com/swopstar/gokit/jobs"
	"github.com/swopstar/gokit/types"
)

type Config struct {
	Server   Server          `toml:"server"`
	Database database.Config `toml:"database"`
	Auth     Auth            `toml:"auth"`
	Jobs     jobs.Config     `toml:"jobs"`
}

type Server struct {
	Network string `toml:"network"`
	Address string `toml:"address"`

	ReadTimeout  types.Duration `toml:"read-timeout"`
	WriteTimeout types.Duration `toml:"write-timeout"`
	IdleTimeout  types.Duration `toml:"idle-timeout"`
}

type Database struct {
	Conn string `toml:"conn"`
}

type Auth struct {
	JWT AuthJWT `toml:"jwt"`
}

type AuthJWT struct {
	KeyPath         string         `toml:"key-path"          comment:"path to private key used to sign JWT tokens"`
	AccessTokenTTL  types.Duration `toml:"access-token-ttl"  comment:"how long an access token should last"`
	RefreshTokenTTL types.Duration `toml:"refresh-token-ttl" comment:"how long a refresh token should last"`
	Issuer          string         `toml:"issuer"            comment:"issuer to use in JWT tokens"`
}

func DefaultConfig(env *config.Env) Config {
	env = env.WithPrefix("SWOPTAPE_")

	return Config{
		Server: Server{
			Network:      env.GetOrDefault("SERVER_NETWORK", "tcp"),
			Address:      env.GetOrDefault("SERVER_ADDRESS", "0.0.0.0:8000"),
			ReadTimeout:  types.Duration{Duration: 15 * time.Second},
			WriteTimeout: types.Duration{Duration: 15 * time.Second},
			IdleTimeout:  types.Duration{Duration: 60 * time.Second},
		},
		Database: database.DefaultConfig(env.WithPrefix("DATABASE_"), "swoptape"),
		Auth: Auth{
			JWT: AuthJWT{
				KeyPath:         "jwt.key",
				AccessTokenTTL:  types.Duration{Duration: 300 * time.Second},
				RefreshTokenTTL: types.Duration{Duration: 90 * 24 * time.Hour},
				Issuer:          "swoptape",
			},
		},
		Jobs: jobs.DefaultConfig(env.WithPrefix("JOBS_")),
	}
}

func LoadConfig(l *slog.Logger, env *config.Env, p string) (cfg Config, err error) {
	configPath := filepath.Join(p, "config.toml")
	l.Info("Loading configuration", "path", configPath)

	file, err := os.Open(configPath)
	if errors.Is(err, os.ErrNotExist) {
		l.Info("Writing new configuration file", "path", configPath)
		if err = DefaultConfig(env).writeConfig(configPath); err != nil {
			err = fmt.Errorf("failed to write configuration file: %w", err)
			return
		}

		return LoadConfig(l, env, p)
	} else if err != nil {
		return
	}
	defer func() { _ = file.Close() }()

	return LoadConfigFromReader(env, file)
}

func (c Config) writeConfig(p string) error {
	file, err := os.Create(p)
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()

	encoder := toml.NewEncoder(file)
	return encoder.Encode(c)
}

func LoadConfigFromReader(env *config.Env, r io.Reader) (cfg Config, err error) {
	cfg = DefaultConfig(env)

	decoder := toml.NewDecoder(r)
	err = decoder.Decode(&cfg)

	// TODO: validate config?

	return
}
