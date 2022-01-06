package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type ConfigCS struct {
	Environment string // develop, staging, production

	CatalogServiceHost string
	CatalogServicePort int

	// context timeout in seconds
	CtxTimeout int

	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func LoadCS() ConfigCS {
	c := ConfigCS{}

	c.Environment = cast.ToString(getOrReturnDefaultCS("ENVIRONMENT", "develop"))

	c.LogLevel = cast.ToString(getOrReturnDefaultCS("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefaultCS("HTTP_PORT", ":8080"))
	c.CatalogServiceHost = cast.ToString(getOrReturnDefaultCS("CATALOG_SERVICE_HOST", "127.0.0.1"))
	c.CatalogServicePort = cast.ToInt(getOrReturnDefaultCS("CATALOG_SERVICE_PORT", 50051))

	c.CtxTimeout = cast.ToInt(getOrReturnDefaultCS("CTX_TIMEOUT", 7))

	return c
}

func getOrReturnDefaultCS(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
