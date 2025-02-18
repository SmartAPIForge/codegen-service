package lib

import (
	"codegen-service/internal/engine/models"
	"fmt"
)

func ComposeProjectId(saf *models.Saf) string {
	return fmt.Sprintf("%s_%s", saf.General.Owner, saf.General.Name)
}
