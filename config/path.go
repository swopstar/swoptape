// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Rareș Nistor

package config

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/swopstar/gokit/config"
)

const EnvPrefix = "SWOPTAPE_"

func DataDir(env *config.Env) (p string) {
	logger := slog.With("group", "config.DataDir")

	p, ok := env.Get("SWOPTAPE_DATA")
	if ok {
		logger.Debug("using env var")
		return
	}

	exePath, err := os.Executable()
	if err != nil {
		p = filepath.Join(exePath, "data")
		_, err := os.Stat(p)
		if err == nil {
			logger.Debug("relative to executable")
			return
		}

		slog.Debug("tried relative to executable, but errored")
	}

	slog.Debug("can't get executable. using relative to cwd", "err", err)
	return "data"
}
