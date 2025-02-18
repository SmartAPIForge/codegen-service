package codegenservice

import (
	"codegen-service/internal/engine"
	"codegen-service/internal/lib"
	"codegen-service/internal/lib/sl"
	"codegen-service/internal/mapper"
	"codegen-service/internal/redis"
	"context"
	"errors"
	codegenProto "github.com/SmartAPIForge/protos/gen/go/codegen"
	"log/slog"
	"math/rand/v2"
)

type CodegenService struct {
	log         *slog.Logger
	redisClient *redis.RedisClient
}

func NewCodegenService(
	log *slog.Logger,
	redisClient *redis.RedisClient,
) *CodegenService {
	return &CodegenService{
		log:         log,
		redisClient: redisClient,
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
	saf.General.Id = lib.ComposeProjectId(saf)
	saf.General.Port = rand.IntN(65000) // TODO - ask deploy-service

	trackingId := lib.NewUUID()
	err = a.redisClient.SetData(trackingId, codegenProto.GenerationStatus_PENDING.String())
	if err != nil {
		log.Error("writing redis error", sl.Err(err))
		return "", err
	}

	log.Info("starting generate:", saf.General.Owner, saf.General.Name)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				err = a.redisClient.SetData(trackingId, codegenProto.GenerationStatus_FAIL.String())
			}
			err = a.redisClient.SetData(trackingId, codegenProto.GenerationStatus_SUCCESS.String())
		}()
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

	return mapper.MapToGenerationStatus(status), nil
}
