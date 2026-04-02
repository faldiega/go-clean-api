package database

import (
	"fmt"
	"go-simple-api/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConfig(cfg *config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMode,
	)

	//dsn := "host=localhost user=postgres password=admin123 dbname=GolangSimpleAPI port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
