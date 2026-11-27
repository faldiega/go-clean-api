package usecase

import (
	"context"
	"time"

	"go-clean-api/internal/domain/entity"
	repo "go-clean-api/internal/domain/repository"
	"go-clean-api/internal/domain/usecase"

	pkgLogger "go-clean-api/pkg/common/logger"

	"go.uber.org/zap"
)

const dbCheckTimeout = 2 * time.Second

type healthUsecase struct {
	repo   repo.HealthRepository
	logger *zap.Logger
}

func NewHealthUsecase(repo repo.HealthRepository, log *zap.Logger) usecase.HealthUsecase {
	return &healthUsecase{repo: repo, logger: log}
}

func (u *healthUsecase) Liveness(_ context.Context) entity.Health {
	return entity.Health{
		Status:    entity.HealthUp,
		Timestamp: time.Now(),
	}
}

func (u *healthUsecase) Readiness(ctx context.Context) entity.Health {
	logger := pkgLogger.FromContext(ctx, u.logger)
	ctx, cancel := context.WithTimeout(ctx, dbCheckTimeout)
	defer cancel()

	start := time.Now()
	dbHealth := entity.ComponentHealth{Status: entity.HealthUp}

	if err := u.repo.PingDatabase(ctx); err != nil {
		// Detail error hanya di log, jangan dikirim ke client.
		logger.Info("health check: database down: %v", zap.Error(err))
		dbHealth.Status = entity.HealthDown
	}
	dbHealth.LatencyMs = time.Since(start).Milliseconds()

	overall := entity.HealthUp
	if dbHealth.Status == entity.HealthDown {
		overall = entity.HealthDown
	}

	return entity.Health{
		Status:    overall,
		Timestamp: time.Now(),
		Components: map[string]entity.ComponentHealth{
			"database": dbHealth,
		},
	}
}
