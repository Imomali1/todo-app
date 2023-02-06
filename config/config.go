package config

import (
	"os"

	"github.com/spf13/cast"
)

type Configs struct {
	HTTP     HTTPConfigs
	Postgres PostgresConfigs
	Redis    RedisConfigs
}

type HTTPConfigs struct {
	Host string
	Port int
}

type RedisConfigs struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type PostgresConfigs struct {
	Host     string
	Port     int
	User     string
	Password string
	DB       string
}

func Load() *Configs {
	cfg := &Configs{}

	cfg.HTTP.Host = cast.ToString(getOrReturnDefault("HTTP_HOST", "localhost"))
	cfg.HTTP.Port = cast.ToInt(getOrReturnDefault("HTTP_PORT", 8080))

	cfg.Redis.Host = cast.ToString(getOrReturnDefault("REDIS_HOST", "cache"))
	cfg.Redis.Port = cast.ToInt(getOrReturnDefault("REDIS_PORT", 6379))
	cfg.Redis.Password = cast.ToString(getOrReturnDefault("REDIS_PASSWORD", ""))
	cfg.Redis.DB = cast.ToInt(getOrReturnDefault("REDIS_DB", 0))

	cfg.Postgres.Host = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "postgres"))
	cfg.Postgres.Port = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))

	return cfg
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
