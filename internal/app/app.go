package app

import (
	grpcapp "codegen-service/internal/app/grpc"
	"codegen-service/internal/config"
	"codegen-service/internal/redis"
	"codegen-service/internal/s3"
	codegenservice "codegen-service/internal/services/codegen"
	packerservice "codegen-service/internal/services/packer"
	"log/slog"
)

type App struct {
	GrpcApp     *grpcapp.GrpcApp
	RedisClient *redis.RedisClient
}

func NewApp(
	log *slog.Logger,
	cfg *config.Config,
) *App {
	redisClient := redis.NewRedisClient(cfg.RedisAddress, cfg.RedisDb)
	s3Client := s3.NewS3Client(cfg)

	packerService := packerservice.NewPackerService(log, s3Client)
	codegenService := codegenservice.NewCodegenService(log, redisClient, packerService)

	grpcApp := grpcapp.NewGrpcApp(
		log,
		codegenService,
		cfg.GRPC.Port,
	)

	return &App{
		GrpcApp:     grpcApp,
		RedisClient: redisClient,
	}
}
