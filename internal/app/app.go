package app

import (
	grpcapp "codegen-service/internal/app/grpc"
	"codegen-service/internal/redis"
	codegenservice "codegen-service/internal/services/codegen"
	"log/slog"
)

type App struct {
	GrpcApp     *grpcapp.GrpcApp
	RedisClient *redis.RedisClient
}

func NewApp(
	log *slog.Logger,
	grpcPort int,
	redisAddress string,
	redisDb int,
) *App {
	redisClient := redis.NewRedisClient(redisAddress, redisDb)
	codegenService := codegenservice.NewCodegenService(log, redisClient)

	grpcApp := grpcapp.NewGrpcApp(
		log,
		codegenService,
		grpcPort,
	)

	return &App{
		GrpcApp:     grpcApp,
		RedisClient: redisClient,
	}
}
