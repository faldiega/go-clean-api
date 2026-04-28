package config

type Config struct {
	AppName       string
	AppSecret     string
	AppHost       string
	AppPort       string
	AppEnv        string
	AppVersion    string
	Database      Database
	Logger        ZapModel
	Jwt           Jwt
	WhitelistURLs []string
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
	LogMaxSize    int
	LogMaxBackups int
	LogMaxAge     int
	LogCompress   bool
}

type Jwt struct {
	JwtSecret  string
	JwtExpired string
}

func LoadConfig() *Config {

	LoadEnv("config.env")

	return &Config{
		AppPort:   Env("APP_PORT", defaultAppPort),
		AppSecret: Env("APP_SECRET", "secret"),
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
			LogMaxSize:    EnvAsInt("LOG_MAX_SIZE", 1),
			LogMaxBackups: EnvAsInt("LOG_MAX_BACKUPS", 3),
			LogMaxAge:     EnvAsInt("LOG_MAX_AGE", 1),
			LogCompress:   EnvAsBool("LOG_COMPRESS", true),
		},
		Jwt: Jwt{
			JwtSecret:  Env("JWT_SECRET", "simplesecret"),
			JwtExpired: Env("JWT_EXPIRE_HOURS", "24"),
		},
		WhitelistURLs: ParseWhitelistURLs(Env("WHITELIST_URLS", "/api/v1/auth")),
	}
}
