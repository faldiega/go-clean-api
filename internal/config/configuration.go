package config

type Config struct {
	App           App
	Database      Database
	Logger        ZapModel
	Jwt           Jwt
	WhitelistURLs []string
	Pagination    Pagination
}

type App struct {
	Name    string
	Secret  string
	Host    string
	Port    string
	Env     string
	Version string
	Timeout int
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
	Secret  string
	Expired int
}

type Pagination struct {
	DefaultPage  int
	DefaultLimit int
}

func Load() *Config {

	LoadEnv("config.env")

	return &Config{
		App: App{
			Port:    Env("APP_PORT", defaultAppPort),
			Secret:  Env("APP_SECRET", "secret"),
			Timeout: EnvAsInt("APP_TIMEOUT", 30),
		},
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
			Secret:  Env("JWT_SECRET", "simplesecret"),
			Expired: EnvAsInt("JWT_EXPIRE_HOURS", 24),
		},
		WhitelistURLs: ParseWhitelistURLs(Env("WHITELIST_URLS", "/api/v1/auth")),
		Pagination: Pagination{
			DefaultPage:  EnvAsInt("PAGINATION_DEFAULT_PAGE", 1),
			DefaultLimit: EnvAsInt("PAGINATION_DEFAULT_LIMIT", 10),
		},
	}
}
