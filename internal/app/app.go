package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"tt2/internal/config"
	"tt2/internal/handler"
	"tt2/internal/middleware"
	"tt2/internal/repository"
	"tt2/internal/repository/postgres"
)

type App struct {
	Cfg *config.Config
	Handler *handler.Handler
}

func New() *App {
	cfg := config.MustLoad()
	log := setupLogger(cfg)
	db := postgres.MustInitDatabase(cfg, log)
	repo := repository.New(db)
	mw := middleware.New(log)
	h := handler.New(log, cfg, mw, repo)
	fmt.Println("successfully done")
	return &App{
		Cfg: cfg,
		Handler: h,
	}
}

func (a *App) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%s", a.Cfg.Port), a.Handler.InitRoutes())
}


func setupLogger(cfg *config.Config) *slog.Logger {
	switch cfg.Env {
	case "local":
		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		logger.Info("Logger initialized")
		return logger
	case "prod":
		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
		logger.Info("Logger initialized")
		return logger
	default:
		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))
		logger.Info("Logger initialized")
		return logger
	}
}