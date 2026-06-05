package main

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	gkconfig "github.com/swopstar/gokit/config"
	"github.com/swopstar/gokit/exitstatus"
	"github.com/swopstar/swoptape/config"
	"github.com/swopstar/swoptape/database"
	"github.com/swopstar/swoptape/services"
	"github.com/swopstar/swoptape/www"
)

var ErrMainExit = errors.New("main exited")
var ErrConnClosed = errors.New("web server closed")

func main() {
	logger := slog.Default()
	env := gkconfig.NewRealEnv()

	dataDir := config.DataDir(env)
	logger.Info("Starting application", "dataDir", dataDir)

	cfg, err := config.LoadConfig(logger, env, dataDir)
	if err != nil {
		logger.Error("Failed to load configuration", "err", err)
		os.Exit(exitstatus.ErrConfig)
	}

	ctx, cancel := context.WithCancelCause(context.Background())

	db, err := database.Open(&cfg.Database, logger.WithGroup("database"))
	if err != nil {
		logger.Error("Failed to open database", "err", err)
		os.Exit(exitstatus.ErrTemp) // TODO: derive exit code from error value
	}

	svc, err := services.New(ctx, &cfg, logger.WithGroup("svc"), db)
	if err != nil {
		logger.Error("Failed to start services", "err", err)
		os.Exit(exitstatus.ErrInternal) // TODO: derive exit code from error value
	}

	srv, err := www.New(ctx, &cfg, logger.WithGroup("www"), svc)
	if err != nil {
		logger.Error("Failed to start web server", "err", err)
		os.Exit(exitstatus.ErrPerm) // TODO: derive exit code from error value
	}

	go func() {
		err := srv.Listen()
		cancel(errors.Join(ErrConnClosed, err))
	}()

	waitForExitSignal()

	logger.Info("Exiting server")
	cancel(ErrMainExit)

	exitStatus := exitstatus.OK

	if err = srv.Close(); err != nil {
		logger.Error("Errors during web server shutdown", "err", err)
		exitStatus = exitstatus.ErrShutdown
	}

	if err = svc.Close(); err != nil {
		logger.Error("Errors during service shutdown", "err", err)
		exitStatus = exitstatus.ErrShutdown
	}

	if err = db.Close(); err != nil {
		logger.Error("Errors during database shutdown", "err", err)
		exitStatus = exitstatus.ErrShutdown
	}

	os.Exit(exitStatus)
}

func waitForExitSignal() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan
}
