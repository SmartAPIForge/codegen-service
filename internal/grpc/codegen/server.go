package authserver

import (
	codegenservice "codegen-service/internal/services/codegen"
	"context"
	"errors"
	codegenProto "github.com/SmartAPIForge/protos/gen/go/codegen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CodegenService interface {
	Generate(
		ctx context.Context,
		saf string,
	) (string, error)
	Track(
		ctx context.Context,
		trackingId string,
	) (codegenProto.GenerationStatus, error)
}

type CodegenServer struct {
	codegenProto.UnimplementedCodegenServiceServer
	codegenService CodegenService
}

func RegisterCodegenServer(
	gRPCServer *grpc.Server,
	codegenService CodegenService,
) {
	codegenProto.RegisterCodegenServiceServer(gRPCServer, &CodegenServer{codegenService: codegenService})
}

func (s *CodegenServer) Generate(
	ctx context.Context,
	in *codegenProto.SafRequest,
) (*codegenProto.TrackDTO, error) {
	if in.Data == "" {
		return nil, status.Error(codes.InvalidArgument, "json is required")
	}

	trackingId, err := s.codegenService.Generate(ctx, in.Data)
	if err != nil {
		if errors.Is(err, codegenservice.ErrInvalidContract) {
			return nil, status.Error(codes.InvalidArgument, "incorrect json provided")
		}
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &codegenProto.TrackDTO{Id: trackingId}, nil
}

func (s *CodegenServer) Track(
	ctx context.Context,
	in *codegenProto.TrackDTO,
) (*codegenProto.GenerationStatusResponse, error) {
	if in.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "tracking id is required")
	}

	statusState, err := s.codegenService.Track(ctx, in.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &codegenProto.GenerationStatusResponse{Status: statusState}, nil
}
