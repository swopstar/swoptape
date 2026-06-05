package www

import (
	"context"
	"log/slog"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swopstar/swoptape/config"
	"github.com/swopstar/swoptape/services"
	"github.com/swopstar/swoptape/www/api"
	"github.com/swopstar/swoptape/www/rpc"
	"github.com/swopstar/swoptape/www/subsonic"
)

type WebServer struct {
	config   *config.Config
	logger   *slog.Logger
	lifetime context.Context

	Router     *gin.Engine
	httpServer *http.Server

	rpc      *rpc.Router
	api      *api.Router
	subsonic *subsonic.Router
}

func New(
	ctx context.Context,
	cfg *config.Config,
	l *slog.Logger,
	svc *services.Services,
) (srv *WebServer, err error) {
	router := gin.Default()

	httpServer := &http.Server{
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout.Duration,
		WriteTimeout: cfg.Server.WriteTimeout.Duration,
		IdleTimeout:  cfg.Server.IdleTimeout.Duration,
	}

	srv = &WebServer{
		lifetime: ctx,
		config:   cfg,
		logger:   l.WithGroup("http"),

		Router:     router,
		httpServer: httpServer,

		rpc:      rpc.New(),
		api:      api.New(),
		subsonic: subsonic.New(),
	}

	srv.addRoutes()
	return
}

func (srv *WebServer) Listen() error {
	listener, err := net.Listen(srv.config.Server.Network, srv.config.Server.Address)
	if err != nil {
		return err
	}

	srv.logger.Info("Starting server", "addr", listener.Addr().String())
	return srv.httpServer.Serve(listener)
}

func (srv *WebServer) Close() error {
	srv.logger.Info("Closing server")
	return srv.httpServer.Close()
}
