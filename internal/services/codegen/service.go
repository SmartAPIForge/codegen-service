package codegenservice

import (
	"codegen-service/internal/engine"
	"codegen-service/internal/lib"
	"codegen-service/internal/lib/sl"
	"context"
	"errors"
	codegenProto "github.com/SmartAPIForge/protos/gen/go/codegen"
	"log/slog"
)

type CodegenService struct {
	log *slog.Logger
}

func NewCodegenService(
	log *slog.Logger,
) *CodegenService {
	return &CodegenService{
		log: log,
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

	log.Info("starting generate:", saf.General.Owner, saf.General.Name)
	go eng.Proceed(saf)

	trackingId := lib.NewUUID()
	// redisClient.set()

	return trackingId, nil
}

func (a *CodegenService) Track(
	_ context.Context,
	trackingId string,
) (codegenProto.GenerationStatus, error) {
	const op = "codegen.Track"

	//log := a.log.With(
	//	slog.String("op", op),
	//	slog.String("id", trackingId),
	//)

	// redisClient.get()

	return codegenProto.GenerationStatus_PENDING, nil
}

func markGenerationAsSuccessful() {

}

func markGenerationAsFailed() {

}
