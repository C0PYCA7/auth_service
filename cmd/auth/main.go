package main

import (
	"auth_service/internal/app"
	"auth_service/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := NewLogger()
	cfg := config.LoadConfig()
	log.Info("Config loaded", slog.Any("config", cfg))

	application := app.New(log, cfg.GRPCconfig.Port, cfg.DbConfig)

	go func() {
		err := application.GRPCServer.Run()
		if err != nil {
			log.Error("Failed to start gRPC server", "error", err)
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	log.Info("Shutting down server...")
	application.GRPCServer.Stop()
	log.Info("Server gracefully stopped")
}

func NewLogger() *slog.Logger {
	var log *slog.Logger

	log = slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
	return log
}
