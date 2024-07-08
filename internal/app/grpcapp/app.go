package grpcapp

import (
	authgrpc "auth_service/internal/grpc/auth"
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int, authService authgrpc.Auth) *App {
	gRPCServer := grpc.NewServer()
	authgrpc.Register(gRPCServer, authService)
	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) Run() error {
	a.log.Info("Starting gRPC Server")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("could not start gRPC Server: %w", err)
	}
	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("could not start gRPC Server: %w", err)
	}
	return nil
}

func (a *App) Stop() {
	a.log.Info("Stopping gRPC Server")
	a.gRPCServer.GracefulStop()
}
