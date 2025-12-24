package server

import (
	"orchid-starter/config"
	"orchid-starter/internal/bootstrap"
	"orchid-starter/internal/bootstrap/server/applications/handler"
	"orchid-starter/middleware"

	"github.com/kataras/iris/v12"
	"github.com/mataharibiz/ward/logging"
)

type Server struct {
	cfg *config.LocalConfig
	app *iris.Application
}

func NewServer(container *bootstrap.Container) *Server {
	srv := &Server{
		cfg: container.GetConfig(),
		app: container.GetApp(),
	}

	// Setup Global middlewares before server initialization
	srv.setupMiddlewares(middleware.SetAPIVersion, middleware.Debug, middleware.Prometheus)

	// Setup routes after server initialization
	srv.setupRoutes(container)

	return srv
}

// setupRoutes configures all application routes
func (s *Server) setupRoutes(container *bootstrap.Container) {
	// Use centralized route management
	handler.SetupAllRoutes(s.app, container)
}

func (s *Server) setupMiddlewares(middlewares ...func(iris.Context)) {
	for _, middleware := range middlewares {
		s.app.Use(middleware)
	}
}

func (s *Server) Run() error {
	// Log version and server info
	cfg := s.cfg
	logging.NewLogger().Info("Version", "version", cfg.AppVersion, "host", cfg.AppHost, "port", cfg.AppPort)

	// Configure application
	app := s.app
	app.Logger().SetLevel(cfg.LogLevel)

	// Start server
	address := cfg.AppHost + ":" + cfg.AppPort
	return app.Run(iris.Addr(address), iris.WithoutServerError(iris.ErrServerClosed))
}
