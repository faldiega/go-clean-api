package repository

import (
	"context"

	repo "go-clean-api/internal/domain/repository"

	"gorm.io/gorm"
)

type healthRepository struct {
	db *gorm.DB
}

func NewHealthRepository(db *gorm.DB) repo.HealthRepository {
	return &healthRepository{db: db}
}

func (r *healthRepository) PingDatabase(ctx context.Context) error {
	if err := r.db.WithContext(ctx); err != nil {
		return err.Error
	}
	// Pastikan DB benar-benar bisa mengeksekusi query, bukan sekadar koneksi terbuka.
	var one int
	return r.db.WithContext(ctx).Raw("SELECT 1").Scan(&one).Error
}
