package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment       string // develop, staging, production
	PostgresHost      string
	PostgresPort      int
	PostgresDatabase  string
	PostgresUser      string
	PostgresPassword  string
	LogLevel          string
	RPCPort           string
	ReviewServiceHost string
	ReviewServicePort int
}

// Load loads environment vars and inflates Config
func Load() Config {
	cfg := Config{}

	cfg.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	cfg.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	cfg.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	cfg.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "catalogdb"))
	cfg.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "jasurbek"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1001"))

	cfg.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))

	cfg.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":9000"))

	return cfg
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
