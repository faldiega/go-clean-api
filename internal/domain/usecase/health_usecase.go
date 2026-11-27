package usecase

import (
	"context"
	"go-clean-api/internal/domain/entity"
)

type HealthUsecase interface {
	// Liveness: hanya memastikan proses hidup, tanpa cek dependency.
	Liveness(ctx context.Context) entity.Health

	// Readiness: memastikan dependency (DB) siap melayani request.
	Readiness(ctx context.Context) entity.Health
}
