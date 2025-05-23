package lib

import codegenProto "github.com/SmartAPIForge/protos/gen/go/codegen"

var generationStatusMap = map[string]codegenProto.GenerationStatus{
	"PENDING": codegenProto.GenerationStatus_PENDING,
	"SUCCESS": codegenProto.GenerationStatus_SUCCESS,
	"FAIL":    codegenProto.GenerationStatus_FAIL,
}

func MapToGenerationStatus(status string) codegenProto.GenerationStatus {
	return generationStatusMap[status]
}

const ExternalStatusPending = "GENERATE_PENDING"
const ExternalStatusGenerated = "GENERATE_SUCCESS"
const ExternalStatusFailedToGenerate = "GENERATE_FAIL"
