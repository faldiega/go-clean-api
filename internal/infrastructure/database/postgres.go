package database

import (
	"fmt"
	"go-simple-api/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConfig(cfg config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Database.DbGolangSimpleApi.DBHost,
		cfg.Database.DbGolangSimpleApi.DBUser,
		cfg.Database.DbGolangSimpleApi.DBPassword,
		cfg.Database.DbGolangSimpleApi.DBName,
		cfg.Database.DbGolangSimpleApi.DBPort,
		cfg.Database.DbGolangSimpleApi.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
