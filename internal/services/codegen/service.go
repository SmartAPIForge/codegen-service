package codegenservice

import (
	"codegen-service/internal/engine"
	"codegen-service/internal/kafka"
	"codegen-service/internal/lib"
	"codegen-service/internal/lib/sl"
	"codegen-service/internal/redis"
	packerservice "codegen-service/internal/services/packer"
	"context"
	"errors"
	codegenProto "github.com/SmartAPIForge/protos/gen/go/codegen"
	"log/slog"
	"math/rand/v2"
)

type CodegenService struct {
	log           *slog.Logger
	redisClient   *redis.RedisClient
	packerService *packerservice.PackerService
	kafkaProducer *kafka.KafkaProducer
}

func NewCodegenService(
	log *slog.Logger,
	redisClient *redis.RedisClient,
	packerService *packerservice.PackerService,
	kafkaProducer *kafka.KafkaProducer,
) *CodegenService {
	return &CodegenService{
		log:           log,
		redisClient:   redisClient,
		packerService: packerService,
		kafkaProducer: kafkaProducer,
	}
}

var (
	ErrInvalidContract = errors.New("invalid contract")
)

func (a *CodegenService) Generate(
	_ context.Context,
	contract string,
) (string, error) {
	const op = "codegen.Generate"

	log := a.log.With(
		slog.String("op", op),
	)

	eng := engine.NewEngine(contract)
	saf, err := eng.ParseSourceToSAF()
	if err != nil {
		log.Error("failed to parse contract", sl.Err(err))
		return "", ErrInvalidContract
	}
	projectId := lib.ComposeProjectId(saf)
	saf.General.Id = projectId
	saf.General.Port = rand.IntN(65000)

	trackingId := lib.NewUUID()
	a.generationPendingStatusFlow(trackingId, projectId)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				a.generationFailStatusFlow(trackingId, projectId)
				return
			}

			url, err := a.packerService.PackAndUpload(projectId)
			if err != nil {
				a.generationFailStatusFlow(trackingId, projectId)
				return
			}

			nativeNewZip := map[string]interface{}{
				"owner": saf.General.Owner,
				"name":  saf.General.Name,
				"url":   url,
			}
			err = a.kafkaProducer.ProduceNewZip(projectId, nativeNewZip)
			if err != nil {
				a.generationFailStatusFlow(trackingId, projectId)
				return
			}

			a.generationSuccessStatusFlow(trackingId, projectId)
		}()

		log.Info("Starting generate:", projectId)
		eng.Proceed(saf)
	}()

	return trackingId, nil
}

func (a *CodegenService) Track(
	_ context.Context,
	trackingId string,
) (codegenProto.GenerationStatus, error) {
	const op = "codegen.Track"

	log := a.log.With(
		slog.String("op", op),
		slog.String("id", trackingId),
	)

	status, err := a.redisClient.GetData(trackingId)
	if err != nil {
		log.Error("reading redis error", sl.Err(err))
		return codegenProto.GenerationStatus_UNKNOWN, err
	}

	return lib.MapToGenerationStatus(status), nil
}

func (a *CodegenService) generationPendingStatusFlow(trackingId, projectId string) {
	a.redisClient.SetData(trackingId, codegenProto.GenerationStatus_PENDING.String(), nil)

	nativeProjectStatus := map[string]interface{}{
		"id":     projectId,
		"status": lib.ExternalStatusPending,
	}
	_ = a.kafkaProducer.ProduceProjectStatus(projectId, nativeProjectStatus)
}

func (a *CodegenService) generationSuccessStatusFlow(trackingId, projectId string) {
	a.redisClient.SetData(trackingId, codegenProto.GenerationStatus_SUCCESS.String(), nil)

	nativeProjectStatus := map[string]interface{}{
		"id":     projectId,
		"status": lib.ExternalStatusGenerated,
	}
	_ = a.kafkaProducer.ProduceProjectStatus(projectId, nativeProjectStatus)
}

func (a *CodegenService) generationFailStatusFlow(trackingId, projectId string) {
	a.redisClient.SetData(trackingId, codegenProto.GenerationStatus_FAIL.String(), nil)

	nativeProjectStatus := map[string]interface{}{
		"id":     projectId,
		"status": lib.ExternalStatusFailedToGenerate,
	}
	_ = a.kafkaProducer.ProduceProjectStatus(projectId, nativeProjectStatus)
}
