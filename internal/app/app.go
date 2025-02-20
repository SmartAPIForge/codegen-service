package app

import (
	grpcapp "codegen-service/internal/app/grpc"
	"codegen-service/internal/config"
	"codegen-service/internal/kafka"
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
	redisClient := redis.NewRedisClient(cfg)
	s3Client := s3.NewS3Client(cfg)
	schemaManager := kafka.NewSchemaManager(cfg)
	kafkaProducer := kafka.NewKafkaProducer(cfg, log, schemaManager)
	packerService := packerservice.NewPackerService(log, s3Client)
	codegenService := codegenservice.NewCodegenService(log, redisClient, packerService, kafkaProducer)

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
