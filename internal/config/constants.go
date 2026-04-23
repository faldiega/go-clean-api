package config

import "time"

// Default constants
const (
	defaultAppName = "GO SIMPLE API"
	defaultAppPort = "5080"

	// default DB
	defaultDBPostgresPort      = 5432
	defaultDBMSSQLPort         = 1433
	defaultDBMaxIdleConnection = 25
	defaultDBMaxIdleTime       = 5 * time.Minute
	defaultDBMaxOpenConnection = 25
	defaultDBMaxLifeConnection = 5 * time.Minute

	// Default HTTP Client Config
	defaultHTTPClientTimeout          = 100 * time.Second
	defaultHTTPClientRetryCount       = 3
	defaultHTTPClientRetryWaitTime    = 10 * time.Second
	defaultHTTPClientRetryMaxWaitTime = 20 * time.Second
)

// Environment constants
const (
	envProd  = "prod"
	envDev   = "dev"
	envLocal = "local"
	envStg   = "stg"
)
