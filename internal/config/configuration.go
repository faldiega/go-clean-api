package config

type Config struct {
	Database Database
	Logger   ZapModel
	Jwt      Jwt
}

type (
	Database struct {
		DbGolangSimpleApi DbConfig
	}

	DbConfig struct {
		DBHost     string
		DBPort     string
		DBUser     string
		DBPassword string
		DBName     string
		DBSSLMode  string
	}
)

type ZapModel struct {
	LogFilename   string
	LogMaxSize    string
	LogMaxBackups string
	LogMaxAge     string
	LogCompress   string
}

type Jwt struct {
	JwtSecret  string
	JwtExpired string
}

func LoadConfig() Config {

	LoadEnv("config.env")

	return Config{
		Database: Database{
			DbGolangSimpleApi: DbConfig{
				DBHost:     Env("DB_HOST", "localhost"),
				DBPort:     Env("DB_PORT", "5432"),
				DBUser:     Env("DB_USER", "postgres"),
				DBPassword: Env("DB_PASSWORD", "postgres"),
				DBName:     Env("DB_NAME", "gosimpleapi"),
				DBSSLMode:  Env("DB_SSLMODE", "disable"),
			},
		},
		Logger: ZapModel{
			LogFilename:   Env("LOG_FILENAME", "app"),
			LogMaxSize:    Env("LOG_MAX_SIZE", "1"),
			LogMaxBackups: Env("LOG_MAX_BACKUPS", "3"),
			LogMaxAge:     Env("LOG_MAX_AGE", "1"),
			LogCompress:   Env("LOG_COMPRESS", "true"),
		},
		Jwt: Jwt{
			JwtSecret:  Env("JWT_SECRET", "simplesecret"),
			JwtExpired: Env("JWT_EXPIRE_HOURS", "24"),
		},
	}
}
