package app

import (
	grpcapp "go_sso/internal/app/grpc"

	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePort string,
	tokenTTL time.Duration,
) *App {
	// TODO: реализовать хранилище
	// TODO: реализовать auth сервис

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
