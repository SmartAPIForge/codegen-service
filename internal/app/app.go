package app

import (
	grpcapp "codegen-service/internal/app/grpc"
	codegenservice "codegen-service/internal/services/codegen"
	"log/slog"
)

type App struct {
	GrpcApp *grpcapp.GrpcApp
}

func NewApp(
	log *slog.Logger,
	grpcPort int,
) *App {
	codegenService := codegenservice.NewCodegenService(log)

	grpcApp := grpcapp.NewGrpcApp(
		log,
		codegenService,
		grpcPort,
	)

	return &App{
		GrpcApp: grpcApp,
	}
}
