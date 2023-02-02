package config

import (
	"github.com/spf13/cast"
	"os"
)

type Configs struct {
	HTTPHost string
	HTTPPort string

	PostgresHost string
	PostgresPort string

	MongoHost string
	MongoPort string
}

func Load() Configs {
	conf := Configs{}

	conf.HTTPHost = cast.ToString(getEnvOrDefault("HTTP_HOST", "localhost"))
	conf.HTTPPort = cast.ToString(getEnvOrDefault("HTTP_PORT", "8080"))

	conf.PostgresHost = cast.ToString(getEnvOrDefault("POSTGRES_HOST", "localhost"))
	conf.PostgresPort = cast.ToString(getEnvOrDefault("POSTGRES_PORT", "5432"))

	conf.MongoHost = cast.ToString(getEnvOrDefault("MONGO_HOST", "localhost"))
	conf.MongoPort = cast.ToString(getEnvOrDefault("MONGO_PORT", "27020"))

	return conf
}

func getEnvOrDefault(key string, def interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if !exists {
		return os.Getenv(key)
	}

	return def
}
