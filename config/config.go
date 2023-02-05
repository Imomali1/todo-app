package config

import (
	"os"

	"github.com/spf13/cast"
)

type Configs struct {
	HTTP            HTTPConfig
	Redis           RedisConfig
	CtxTimeout      int
	UserServicePath string
	LogLevel        string
}

type HTTPConfig struct {
	Host string
	Port string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func Load() Configs {
	cfg := Configs{}

	cfg.HTTP.Host = cast.ToString(getOrReturnDefault("HTTP_HOST", "localhost"))
	cfg.HTTP.Port = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8050"))

	cfg.Redis.Host = cast.ToString(getOrReturnDefault("REDIS_HOST", "cache"))
	cfg.Redis.Port = cast.ToString(getOrReturnDefault("REDIS_PORT", ":6379"))
	cfg.Redis.Password = cast.ToString(getOrReturnDefault("REDIS_PASSWORD", ""))
	cfg.CtxTimeout = cast.ToInt(getOrReturnDefault("REDIS_DB", 0))

	return cfg
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
