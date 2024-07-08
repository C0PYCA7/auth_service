package app

import (
	"auth_service/internal/app/grpcapp"
	"auth_service/internal/config"
	"auth_service/internal/services/auth"
	"auth_service/internal/storage/postgres"

	"log/slog"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(log *slog.Logger, port int, cfg config.Database) *App {

	storage, err := postgres.New(cfg)
	if err != nil {
		log.Error("Failed to initialize database", "error", err)
	}

	authService := auth.New(log, storage)

	grpcApp := grpcapp.New(log, port, authService)
	return &App{
		GRPCServer: grpcApp,
	}
}
